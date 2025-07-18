package ui

import (
	"fmt"
	"strings"
	"time"

	"syntaxrush/core"
	"syntaxrush/theme"

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

	// Metrics
	metrics *core.Metrics
	timer   *core.Timer

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

// TickMsg is sent every second to update metrics
type TickMsg time.Time

// NewModel creates a new application model
func NewModel() *Model {
	parser := core.NewParser()
	metrics := core.NewMetrics()
	timer := core.NewTimer()

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
		parser:       parser,
		metrics:      metrics,
		timer:        timer,
		theme:        theme.NewDarkTheme(),
		state:        StateWelcome,
		maxViewLines: 20,
		filename:     "sample.go",
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
	m.viewportStart = 0
	m.sessionComplete = false
	m.timer.Reset()
	m.metrics.Reset()
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
	default:
		if len(msg.String()) == 1 {
			m.userInput += msg.String()
			m.currentPos = len(m.userInput)

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

	// Calculate accuracy for this line
	m.metrics.AddLine(m.userInput, currentCode)

	// Move to next line
	m.currentLine++
	m.userInput = ""
	m.currentPos = 0

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

// getCurrentLine returns the current line to type
func (m *Model) getCurrentLine() string {
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

// Helper methods for rendering will be implemented in view.go
