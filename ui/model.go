package ui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/vamshi1188/SyntaxRush/core"
	"github.com/vamshi1188/SyntaxRush/theme"

	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the application state
type Model struct {
	// File content and parsing
	parser      *core.Parser
	codeLines   []string
	currentLine int
	totalLines  int

	// Typing state
	userInput      string
	correctChars   int
	incorrectChars int
	currentPos     int
	lastMistakePos int // Track last mistake position to avoid repeated sounds

	// Typing history for completed lines
	completedLines map[int]string // Maps line number to user's typed input

	// Metrics
	metrics *core.Metrics
	timer   *core.Timer
	audio   *core.AudioManager
	mpi     *core.MusclePowerIndicator // Muscle Power Indicator

	// UI state
	width         int
	height        int
	theme         *theme.Theme
	viewportStart int
	maxViewLines  int

	// App state
	state    AppState
	message  string
	filename string
	quitting bool

	// File input state
	fileInput string
	fileError string

	// Session summary
	sessionComplete bool
	finalStats      core.SessionStats
}

type AppState int

const (
	StateWelcome AppState = iota
	StateTyping
	StateSummary
	StateFileSelect
)

// playErrorSound plays a beep sound when a mistake is made
// Uses the oto audio library for high-quality cross-platform sound
func (m *Model) playErrorSound() {
	// Play audio if available
	if m.audio != nil {
		m.audio.PlayErrorBeep()
	} else {
		// Fallback to terminal bell if audio initialization failed
		os.Stdout.Write([]byte("\a"))
		os.Stdout.Sync()
	}
}

// TickMsg is sent every second to update metrics
type TickMsg time.Time

// NewModel creates a new application model
func NewModel() *Model {
	parser := core.NewParser()
	metrics := core.NewMetrics()
	timer := core.NewTimer()
	mpi := core.NewMusclePowerIndicator()

	// Initialize audio manager (gracefully handle errors)
	audio, err := core.NewAudioManager()
	if err != nil {
		// If audio fails to initialize, continue without sound
		audio = nil
	}

	// Load default sample code
	sampleCode := `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    
    numbers := []int{1, 2, 3, 4, 5}
    
    for i, num := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", i, num)
    }
    
    result := sum(numbers)
    fmt.Printf("Sum: %d\n", result)
}

func sum(nums []int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}`

	model := &Model{
		parser:         parser,
		metrics:        metrics,
		timer:          timer,
		audio:          audio, // Add audio manager
		mpi:            mpi,   // Add muscle power indicator
		theme:          theme.NewDarkTheme(),
		state:          StateWelcome,
		maxViewLines:   20,
		filename:       "sample.go",
		lastMistakePos: -1,                   // Initialize mistake tracking
		completedLines: make(map[int]string), // Initialize typing history
	}

	model.codeLines = strings.Split(sampleCode, "\n")
	model.totalLines = len(model.codeLines)

	return model
}

// LoadFile loads a code file for typing practice
func (m *Model) LoadFile(filepath string) error {
	content, err := m.parser.ParseFile(filepath)
	if err != nil {
		return err
	}

	if content == "" {
		return fmt.Errorf("file is empty")
	}

	m.codeLines = strings.Split(content, "\n")
	m.totalLines = len(m.codeLines)

	// Extract just the filename for display
	if lastSlash := strings.LastIndex(filepath, "/"); lastSlash != -1 {
		m.filename = filepath[lastSlash+1:]
	} else {
		m.filename = filepath
	}

	m.resetSession()

	return nil
}

// expandFilePath expands shortcuts and handles relative paths
func (m *Model) expandFilePath(input string) string {
	// Handle built-in sample shortcuts
	switch strings.ToLower(input) {
	case "go", "sample.go":
		return "assets/sample.go"
	case "py", "python", "sample.py":
		return "assets/sample.py"
	case "js", "javascript", "sample.js":
		return "assets/sample.js"
	case "cpp", "c++", "sample.cpp":
		return "assets/sample.cpp"
	default:
		// Return the input as-is for regular file paths
		return input
	}
}

// resetSession resets the typing session
func (m *Model) resetSession() {
	m.currentLine = 0
	m.userInput = ""
	m.correctChars = 0
	m.incorrectChars = 0
	m.currentPos = 0
	m.lastMistakePos = -1 // Reset mistake tracking
	m.viewportStart = 0
	m.sessionComplete = false
	m.completedLines = make(map[int]string) // Reset typing history
	m.timer.Reset()
	m.metrics.Reset()
	m.mpi.Reset() // Reset muscle power indicator
	m.state = StateTyping
}

// Init implements tea.Model
func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		tickCmd(),
	)
}

// tickCmd returns a command that sends TickMsg every second
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

// Update implements tea.Model
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.maxViewLines = m.height - 10 // Reserve space for UI elements

	case tea.KeyMsg:
		return m.handleKeyPress(msg)

	case TickMsg:
		if m.state == StateTyping && m.timer.IsRunning() {
			m.updateMetrics()
		}
		return m, tickCmd()
	}

	return m, nil
}

// handleKeyPress handles keyboard input
func (m *Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.state {
	case StateWelcome:
		return m.handleWelcomeKeys(msg)
	case StateTyping:
		return m.handleTypingKeys(msg)
	case StateSummary:
		return m.handleSummaryKeys(msg)
	case StateFileSelect:
		return m.handleFileSelectKeys(msg)
	}
	return m, nil
}

// handleWelcomeKeys handles keys in welcome state
func (m *Model) handleWelcomeKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q", "esc":
		m.quitting = true
		return m, tea.Quit
	case "ctrl+u":
		m.state = StateFileSelect
		m.message = "Enter file path (or press Enter for sample): "
		m.fileInput = ""
		m.fileError = ""
	case "enter", " ":
		m.resetSession()
		m.timer.Start()
	}
	return m, nil
}

// handleTypingKeys handles keys during typing practice
func (m *Model) handleTypingKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "esc":
		m.state = StateWelcome
		m.timer.Stop()
		return m, nil
	case "ctrl+u":
		m.state = StateFileSelect
		m.message = "Enter file path: "
		m.fileInput = ""
		m.fileError = ""
		return m, nil
	case "ctrl+r":
		m.resetSession()
		m.timer.Start()
		return m, nil
	case "enter":
		return m.handleLineComplete(), nil
	case "backspace":
		// Backspace is disabled during typing practice, but track the attempt
		m.mpi.RecordKeystroke(0, false, true)
		// Play error sound to indicate backspace is not allowed
		m.playErrorSound()
		return m, nil
	default:
		if len(msg.String()) == 1 {
			oldInputLen := len(m.userInput)
			m.userInput += msg.String()
			m.currentPos = len(m.userInput)

			// Check if this character is a mistake
			currentLine := m.getCurrentLine()
			isCorrect := false

			if oldInputLen < len(currentLine) {
				expectedChar := rune(currentLine[oldInputLen])
				typedChar := rune(msg.String()[0])
				isCorrect = expectedChar == typedChar

				// If this is a new mistake (not repeating at same position)
				if !isCorrect && m.lastMistakePos != oldInputLen {
					m.playErrorSound()
					m.lastMistakePos = oldInputLen
				}
			} else {
				// Typing beyond the line length is also a mistake
				isCorrect = false
				if m.lastMistakePos != oldInputLen {
					m.playErrorSound()
					m.lastMistakePos = oldInputLen
				}
			}

			// Record keystroke in MPI
			m.mpi.RecordKeystroke(rune(msg.String()[0]), isCorrect, false)

			// Start timer on first keypress
			if !m.timer.IsRunning() {
				m.timer.Start()
			}
		}
	}

	m.updateMetrics()
	return m, nil
}

// handleLineComplete processes when user presses Enter
func (m *Model) handleLineComplete() *Model {
	currentCode := m.getCurrentLine()

	// Store the user's input for this completed line
	m.completedLines[m.currentLine] = m.userInput

	// Calculate accuracy for this line
	m.metrics.AddLine(m.userInput, currentCode)

	// Play success sound if line was typed correctly
	if m.userInput == currentCode && m.audio != nil {
		m.audio.PlaySuccessSound()
	}

	// Move to next line
	m.currentLine++
	m.userInput = ""
	m.currentPos = 0
	m.lastMistakePos = -1 // Reset mistake tracking for new line

	// Check if we've completed all lines
	if m.currentLine >= m.totalLines {
		m.completeSession()
	} else {
		// Update viewport if needed
		m.updateViewport()
	}

	return m
}

// handleSummaryKeys handles keys in session summary
func (m *Model) handleSummaryKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q", "esc":
		m.quitting = true
		return m, tea.Quit
	case "r":
		m.resetSession()
		m.timer.Start()
	case "u":
		m.state = StateFileSelect
		m.message = "Enter file path: "
		m.fileInput = ""
		m.fileError = ""
	case "enter", " ":
		m.state = StateWelcome
	}
	return m, nil
}

// handleFileSelectKeys handles file selection input
func (m *Model) handleFileSelectKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "esc":
		m.state = StateWelcome
		m.message = ""
		m.fileInput = ""
		m.fileError = ""
	case "enter":
		if m.fileInput == "" {
			// Empty input - use sample file
			m.state = StateWelcome
			m.message = ""
			m.fileError = ""
		} else {
			// Try to load the specified file
			filePath := m.expandFilePath(m.fileInput)
			err := m.LoadFile(filePath)
			if err != nil {
				m.fileError = err.Error()
			} else {
				// Successfully loaded file
				m.state = StateWelcome
				m.message = "File loaded successfully: " + m.filename
				m.fileInput = ""
				m.fileError = ""
			}
		}
	case "backspace":
		if len(m.fileInput) > 0 {
			m.fileInput = m.fileInput[:len(m.fileInput)-1]
			m.fileError = "" // Clear error when user starts typing
		}
	default:
		// Add printable characters to file input
		if len(msg.String()) == 1 && msg.String()[0] >= 32 && msg.String()[0] <= 126 {
			m.fileInput += msg.String()
			m.fileError = "" // Clear error when user starts typing
		}
	}
	return m, nil
}

// getCurrentLine returns the current line to type (trimmed of leading whitespace)
func (m *Model) getCurrentLine() string {
	if m.currentLine >= len(m.codeLines) {
		return ""
	}
	return strings.TrimLeft(m.codeLines[m.currentLine], " \t")
}

// getCurrentLineRaw returns the current line with original whitespace (for display)
func (m *Model) getCurrentLineRaw() string {
	if m.currentLine >= len(m.codeLines) {
		return ""
	}
	return m.codeLines[m.currentLine]
}

// updateViewport updates the viewport to keep current line visible
func (m *Model) updateViewport() {
	if m.currentLine >= m.viewportStart+m.maxViewLines {
		m.viewportStart = m.currentLine - m.maxViewLines + 1
	}
	if m.currentLine < m.viewportStart {
		m.viewportStart = m.currentLine
	}
	if m.viewportStart < 0 {
		m.viewportStart = 0
	}
}

// updateMetrics updates real-time typing metrics
func (m *Model) updateMetrics() {
	if m.timer.IsRunning() {
		elapsed := m.timer.Elapsed()
		m.metrics.UpdateRealTime(m.userInput, m.getCurrentLine(), elapsed)
	}
}

// completeSession finishes the typing session
func (m *Model) completeSession() {
	m.timer.Stop()
	m.sessionComplete = true
	m.state = StateSummary

	// Calculate final statistics
	m.finalStats = m.metrics.GetSessionStats(m.timer.Elapsed())
}

// View implements tea.Model
func (m *Model) View() string {
	if m.quitting {
		return "Thanks for using CodeType!\n"
	}

	switch m.state {
	case StateWelcome:
		return m.renderWelcome()
	case StateTyping:
		return m.renderTyping()
	case StateSummary:
		return m.renderSummary()
	case StateFileSelect:
		return m.renderFileSelect()
	}

	return ""
}

// Cleanup properly closes audio resources
func (m *Model) Cleanup() {
	if m.audio != nil {
		m.audio.Close()
	}
}

// SetAudioEnabled enables or disables audio feedback
func (m *Model) SetAudioEnabled(enabled bool) {
	if !enabled && m.audio != nil {
		m.audio.Close()
		m.audio = nil
	}
}

// StartPracticeDirectly skips welcome screen and starts practice immediately
func (m *Model) StartPracticeDirectly() {
	m.resetSession()
	m.timer.Start()
}

// GetFinalStats returns the final session statistics
func (m *Model) GetFinalStats() core.SessionStats {
	if m.sessionComplete {
		return m.finalStats
	}
	// Return current stats if session is ongoing
	return m.metrics.GetSessionStats(m.timer.Elapsed())
}

// GetMPIStats returns muscle power indicator statistics
func (m *Model) GetMPIStats() map[string]interface{} {
	return m.mpi.GetStats()
}

// Helper methods for rendering will be implemented in view.go
