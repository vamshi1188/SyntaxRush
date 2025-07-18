package core

import (
	"time"
	"unicode"
)

// Metrics handles calculation of typing statistics
type Metrics struct {
	totalCharacters   int
	correctCharacters int
	mistakes          int
	totalWords        int
	startTime         time.Time
	lines             []LineStats
	realTimeStats     RealTimeStats
}

// LineStats stores statistics for individual lines
type LineStats struct {
	Original  string
	UserInput string
	Mistakes  int
	Accuracy  float64
	TimeSpent time.Duration
	CharCount int
}

// RealTimeStats stores current typing statistics
type RealTimeStats struct {
	WPM         float64
	CPM         float64
	Accuracy    float64
	Mistakes    int
	ElapsedTime time.Duration
}

// SessionStats stores final session statistics
type SessionStats struct {
	TotalTime       time.Duration
	WPM             float64
	CPM             float64
	Accuracy        float64
	TotalMistakes   int
	TotalCharacters int
	LinesCompleted  int
	ErrorHeatmap    map[int]int // Position -> mistake count
}

// NewMetrics creates a new metrics instance
func NewMetrics() *Metrics {
	return &Metrics{
		lines:         make([]LineStats, 0),
		realTimeStats: RealTimeStats{},
	}
}

// Reset resets all metrics
func (m *Metrics) Reset() {
	m.totalCharacters = 0
	m.correctCharacters = 0
	m.mistakes = 0
	m.totalWords = 0
	m.lines = make([]LineStats, 0)
	m.realTimeStats = RealTimeStats{}
}

// AddLine adds statistics for a completed line
func (m *Metrics) AddLine(userInput, original string) {
	lineStats := m.calculateLineStats(userInput, original)
	m.lines = append(m.lines, lineStats)

	// Update totals
	m.totalCharacters += lineStats.CharCount
	m.correctCharacters += (lineStats.CharCount - lineStats.Mistakes)
	m.mistakes += lineStats.Mistakes
	m.totalWords += m.countWords(original)
}

// calculateLineStats calculates statistics for a single line
func (m *Metrics) calculateLineStats(userInput, original string) LineStats {
	mistakes := 0
	charCount := len(original)

	// Count character mismatches
	minLen := min(len(userInput), len(original))
	for i := 0; i < minLen; i++ {
		if userInput[i] != original[i] {
			mistakes++
		}
	}

	// Count extra or missing characters
	if len(userInput) != len(original) {
		mistakes += abs(len(userInput) - len(original))
	}

	accuracy := 100.0
	if charCount > 0 {
		accuracy = float64(charCount-mistakes) / float64(charCount) * 100
		if accuracy < 0 {
			accuracy = 0
		}
	}

	return LineStats{
		Original:  original,
		UserInput: userInput,
		Mistakes:  mistakes,
		Accuracy:  accuracy,
		CharCount: charCount,
	}
}

// UpdateRealTime updates real-time statistics during typing
func (m *Metrics) UpdateRealTime(currentInput, currentLine string, elapsed time.Duration) {
	if elapsed == 0 {
		return
	}

	// Calculate current line mistakes
	currentMistakes := 0
	minLen := min(len(currentInput), len(currentLine))
	for i := 0; i < minLen; i++ {
		if currentInput[i] != currentLine[i] {
			currentMistakes++
		}
	}

	// Add extra characters as mistakes
	if len(currentInput) > len(currentLine) {
		currentMistakes += len(currentInput) - len(currentLine)
	}

	// Calculate total characters typed (including current line)
	totalTyped := m.totalCharacters + len(currentInput)

	// Calculate total mistakes (including current line)
	totalMistakes := m.mistakes + currentMistakes

	// Calculate accuracy
	accuracy := 100.0
	if totalTyped > 0 {
		correct := totalTyped - totalMistakes
		accuracy = float64(correct) / float64(totalTyped) * 100
		if accuracy < 0 {
			accuracy = 0
		}
	}

	// Calculate WPM and CPM
	minutes := elapsed.Minutes()
	if minutes > 0 {
		// Use total words typed (completed lines + current line words)
		currentWords := m.countWords(currentInput)
		totalWords := m.totalWords + currentWords

		wpm := float64(totalWords) / minutes
		cpm := float64(totalTyped) / minutes

		m.realTimeStats = RealTimeStats{
			WPM:         wpm,
			CPM:         cpm,
			Accuracy:    accuracy,
			Mistakes:    totalMistakes,
			ElapsedTime: elapsed,
		}
	}
}

// GetCurrentStats returns current real-time statistics
func (m *Metrics) GetCurrentStats() RealTimeStats {
	return m.realTimeStats
}

// GetSessionStats calculates final session statistics
func (m *Metrics) GetSessionStats(totalTime time.Duration) SessionStats {
	totalChars := m.totalCharacters
	totalMistakes := m.mistakes

	accuracy := 100.0
	if totalChars > 0 {
		correct := totalChars - totalMistakes
		accuracy = float64(correct) / float64(totalChars) * 100
		if accuracy < 0 {
			accuracy = 0
		}
	}

	var wpm, cpm float64
	if totalTime.Minutes() > 0 {
		wpm = float64(m.totalWords) / totalTime.Minutes()
		cpm = float64(totalChars) / totalTime.Minutes()
	}

	// Generate error heatmap
	heatmap := make(map[int]int)
	for _, line := range m.lines {
		// This is simplified - in a full implementation,
		// you'd track position-specific errors
		if line.Mistakes > 0 {
			heatmap[len(line.Original)] += line.Mistakes
		}
	}

	return SessionStats{
		TotalTime:       totalTime,
		WPM:             wpm,
		CPM:             cpm,
		Accuracy:        accuracy,
		TotalMistakes:   totalMistakes,
		TotalCharacters: totalChars,
		LinesCompleted:  len(m.lines),
		ErrorHeatmap:    heatmap,
	}
}

// countWords counts the number of words in a string
func (m *Metrics) countWords(text string) int {
	if len(text) == 0 {
		return 0
	}

	words := 0
	inWord := false

	for _, char := range text {
		if unicode.IsSpace(char) {
			inWord = false
		} else if !inWord {
			words++
			inWord = true
		}
	}

	return words
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
