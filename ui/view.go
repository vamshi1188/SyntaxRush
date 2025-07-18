package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// renderWelcome renders the welcome screen
func (m *Model) renderWelcome() string {
	instructions := []string{
		"",
		"🚀 SyntaxRush - Welcome! 🚀",
		"🧠⚡ Improve your coding speed and accuracy.",
		"",
		"Practice typing real code. Master syntax.",
		"🔊 Audio feedback for mistakes and success!",
		"",
		"📁 Current file: " + m.filename,
		fmt.Sprintf("📄 Lines: %d", m.totalLines),
		"",
		"Controls:",
		"  Enter/Space - Start typing practice",
		"  Ctrl+U      - Upload new file",
		"  Q/Esc       - Quit",
		"",
		"Press Enter or Space to begin!",
	}

	content := strings.Join(instructions, "\n")
	styledContent := m.theme.Text.Render(content)

	if m.message != "" {
		message := m.theme.Error.Render(m.message)
		return lipgloss.JoinVertical(lipgloss.Center, styledContent, "", message)
	}

	return styledContent
}

// renderTyping renders the main typing interface
func (m *Model) renderTyping() string {
	// Header with file info
	header := m.renderHeader()

	// Code display pane
	codePane := m.renderCodePane()

	// Typing input pane
	inputPane := m.renderInputPane()

	// Metrics panel
	metricsPanel := m.renderMetrics()

	// Controls help
	controls := m.renderControls()

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		"",
		codePane,
		"",
		inputPane,
		"",
		metricsPanel,
		"",
		controls,
	)
}

// renderHeader renders the file information header
func (m *Model) renderHeader() string {
	progress := fmt.Sprintf("%d/%d", m.currentLine+1, m.totalLines)
	percentage := float64(m.currentLine) / float64(m.totalLines) * 100

	title := fmt.Sprintf("📁 %s", m.filename)
	progressInfo := fmt.Sprintf("Progress: %s (%.1f%%)", progress, percentage)

	headerStyle := m.theme.Header.Width(m.width - 2)
	return headerStyle.Render(fmt.Sprintf("%s • %s", title, progressInfo))
}

// renderCodePane renders the code display area
func (m *Model) renderCodePane() string {
	var lines []string

	// Calculate visible range
	startLine := m.viewportStart
	endLine := startLine + m.maxViewLines
	if endLine > m.totalLines {
		endLine = m.totalLines
	}

	for i := startLine; i < endLine; i++ {
		lineNum := fmt.Sprintf("%3d", i+1)
		code := m.codeLines[i]

		if i == m.currentLine {
			// Highlight current line
			styledLine := m.theme.CurrentLine.Render(fmt.Sprintf("%s │ %s", lineNum, code))
			lines = append(lines, styledLine)
		} else {
			styledLine := m.theme.CodeLine.Render(fmt.Sprintf("%s │ %s", lineNum, code))
			lines = append(lines, styledLine)
		}
	}

	content := strings.Join(lines, "\n")

	title := m.theme.PaneTitle.Render("📖 Code to Type")
	paneStyle := m.theme.CodePane.Width(m.width - 4).Height(m.maxViewLines + 2)

	return lipgloss.JoinVertical(lipgloss.Left, title, paneStyle.Render(content))
}

// renderInputPane renders the typing input area
func (m *Model) renderInputPane() string {
	currentCode := m.getCurrentLine()

	// Create styled input showing correct/incorrect characters
	var styledInput strings.Builder

	for i, char := range currentCode {
		if i < len(m.userInput) {
			userChar := rune(m.userInput[i])
			if userChar == char {
				// Correct character
				styledInput.WriteString(m.theme.CorrectChar.Render(string(char)))
			} else {
				// Incorrect character
				styledInput.WriteString(m.theme.IncorrectChar.Render(string(char)))
			}
		} else if i == len(m.userInput) {
			// Current cursor position
			styledInput.WriteString(m.theme.Cursor.Render(string(char)))
		} else {
			// Remaining characters
			styledInput.WriteString(m.theme.RemainingChar.Render(string(char)))
		}
	}

	// Show extra characters if user typed too much
	if len(m.userInput) > len(currentCode) {
		extra := m.userInput[len(currentCode):]
		styledInput.WriteString(m.theme.ExtraChar.Render(extra))
	}

	title := m.theme.PaneTitle.Render("⌨️  Your Input")

	// Create the input display
	inputDisplay := styledInput.String()
	if inputDisplay == "" {
		inputDisplay = m.theme.Cursor.Render("█") // Show cursor when no input
	}

	paneStyle := m.theme.InputPane.Width(m.width - 4).Height(3)

	return lipgloss.JoinVertical(lipgloss.Left, title, paneStyle.Render(inputDisplay))
}

// renderMetrics renders the real-time metrics panel
func (m *Model) renderMetrics() string {
	if !m.timer.IsRunning() && m.timer.Elapsed() == 0 {
		// Show default metrics before starting
		metrics := []string{
			"⏱️  Time: 00:00",
			"🎯 Accuracy: --%",
			"⚡ WPM: --",
			"📊 CPM: --",
			"❌ Mistakes: 0",
		}
		content := strings.Join(metrics, " │ ")
		return m.theme.MetricsPanel.Width(m.width - 2).Render(content)
	}

	stats := m.metrics.GetCurrentStats()
	elapsed := m.timer.Elapsed()

	timeStr := formatDuration(elapsed)

	metrics := []string{
		fmt.Sprintf("⏱️  Time: %s", timeStr),
		fmt.Sprintf("🎯 Accuracy: %.1f%%", stats.Accuracy),
		fmt.Sprintf("⚡ WPM: %.0f", stats.WPM),
		fmt.Sprintf("📊 CPM: %.0f", stats.CPM),
		fmt.Sprintf("❌ Mistakes: %d", stats.Mistakes),
	}

	content := strings.Join(metrics, " │ ")
	return m.theme.MetricsPanel.Width(m.width - 2).Render(content)
}

// renderControls renders the control help
func (m *Model) renderControls() string {
	controls := "Ctrl+R: Retry │ Ctrl+U: Upload │ Esc: Menu"
	return m.theme.Controls.Render(controls)
}

// renderSummary renders the session completion summary
func (m *Model) renderSummary() string {
	title := m.theme.Title.Render("🎉 SyntaxRush Session Complete!")

	stats := []string{
		fmt.Sprintf("📁 File: %s", m.filename),
		fmt.Sprintf("📄 Lines completed: %d", m.totalLines),
		fmt.Sprintf("⏱️  Total time: %s", formatDuration(m.finalStats.TotalTime)),
		fmt.Sprintf("🎯 Final accuracy: %.1f%%", m.finalStats.Accuracy),
		fmt.Sprintf("⚡ Average WPM: %.1f", m.finalStats.WPM),
		fmt.Sprintf("📊 Average CPM: %.1f", m.finalStats.CPM),
		fmt.Sprintf("❌ Total mistakes: %d", m.finalStats.TotalMistakes),
		"",
		"🏆 Great job! Keep practicing to improve your speed and accuracy.",
	}

	statsContent := strings.Join(stats, "\n")
	styledStats := m.theme.Summary.Render(statsContent)

	controls := []string{
		"",
		"What's next?",
		"  R - Retry this file",
		"  U - Upload new file",
		"  Enter/Space - Back to menu",
		"  Q/Esc - Quit",
	}

	controlsContent := strings.Join(controls, "\n")
	styledControls := m.theme.Text.Render(controlsContent)

	return lipgloss.JoinVertical(lipgloss.Center, title, "", styledStats, styledControls)
}

// renderFileSelect renders the file selection screen
func (m *Model) renderFileSelect() string {
	title := m.theme.Title.Render("📁 SyntaxRush - File Upload")

	content := []string{
		"Load a code file for typing practice",
		"",
		"Quick shortcuts:",
		"  • 'go' or 'sample.go' - Go calculator example",
		"  • 'py' or 'sample.py' - Python data processor",
		"  • 'js' or 'sample.js' - JavaScript task manager",
		"  • 'cpp' or 'sample.cpp' - C++ grade system",
		"",
		"Or enter full/relative path to your file:",
		"  • ./mycode.go",
		"  • /home/user/project/main.py",
		"  • ../examples/demo.js",
		"",
		"Supported formats:",
		"  .go .py .js .ts .jsx .tsx .cpp .c .java .rs",
		"",
	}

	// Show the input prompt
	prompt := "Enter file path: "

	// Show the current input with cursor
	inputDisplay := m.fileInput + "█"
	if m.fileInput == "" {
		inputDisplay = "█"
	}

	styledPrompt := m.theme.Text.Render(prompt)
	styledInput := m.theme.InputPane.Render(inputDisplay)

	content = append(content, "")

	// Show error if any
	if m.fileError != "" {
		errorMsg := "❌ Error: " + m.fileError
		content = append(content, errorMsg)
		content = append(content, "")
	}

	content = append(content, "Controls:")
	content = append(content, "  Enter - Load file (or use sample if empty)")
	content = append(content, "  Backspace - Delete character")
	content = append(content, "  Esc - Back to menu")

	styledContent := m.theme.Text.Render(strings.Join(content, "\n"))

	inputSection := lipgloss.JoinHorizontal(lipgloss.Left, styledPrompt, styledInput)

	return lipgloss.JoinVertical(lipgloss.Center, title, "", styledContent, "", inputSection)
}

// formatDuration formats a duration as MM:SS
func formatDuration(d time.Duration) string {
	minutes := int(d.Minutes())
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
