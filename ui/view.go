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

	// Combined code display with typing
	codePane := m.renderUnifiedCodePane()

	// Muscle Power Indicator
	mpiPanel := m.renderMusclePowerIndicator()

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
		mpiPanel,
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
	currentCodeRaw := m.getCurrentLineRaw()                 // Original line with indentation
	currentCode := m.getCurrentLine()                       // Trimmed line for typing
	leadingSpaces := len(currentCodeRaw) - len(currentCode) // Calculate indentation

	// Create styled input showing correct/incorrect characters
	var styledInput strings.Builder

	// First, add the leading spaces as already "typed" (in gray/dim style)
	if leadingSpaces > 0 {
		indentation := currentCodeRaw[:leadingSpaces]
		styledInput.WriteString(m.theme.RemainingChar.Render(indentation))
	}

	// Then handle the actual typing part (trimmed content)
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

// renderUnifiedCodePane renders the code with integrated typing display
func (m *Model) renderUnifiedCodePane() string {
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
			// This is the current line being typed - show typing progress
			styledLine := m.renderCurrentLineWithTyping(lineNum, code)
			lines = append(lines, styledLine)
		} else if userInput, isCompleted := m.completedLines[i]; isCompleted {
			// This line was completed - show it with color coding
			styledLine := m.renderCompletedLineWithColors(lineNum, code, userInput)
			lines = append(lines, styledLine)
		} else {
			// Regular line display (not yet reached)
			styledLine := m.theme.CodeLine.Render(fmt.Sprintf("%s │ %s", lineNum, code))
			lines = append(lines, styledLine)
		}
	}

	content := strings.Join(lines, "\n")

	title := m.theme.PaneTitle.Render("📖 Code Practice")
	paneStyle := m.theme.CodePane.Width(m.width - 4).Height(m.maxViewLines + 2)

	return lipgloss.JoinVertical(lipgloss.Left, title, paneStyle.Render(content))
}

// renderCurrentLineWithTyping renders the current line with typing progress
func (m *Model) renderCurrentLineWithTyping(lineNum, codeLine string) string {
	currentCodeRaw := codeLine                              // Original line with indentation
	currentCode := m.getCurrentLine()                       // Trimmed line for typing
	leadingSpaces := len(currentCodeRaw) - len(currentCode) // Calculate indentation

	// Build the display line: line number + separator + styled content
	var lineBuilder strings.Builder
	lineBuilder.WriteString(lineNum)
	lineBuilder.WriteString(" │ ")

	// Add the leading spaces (show them as dim/gray)
	if leadingSpaces > 0 {
		indentation := currentCodeRaw[:leadingSpaces]
		lineBuilder.WriteString(m.theme.RemainingChar.Render(indentation))
	}

	// Now handle the actual typing content
	for i, char := range currentCode {
		if i < len(m.userInput) {
			userChar := rune(m.userInput[i])
			if userChar == char {
				// Correct character
				lineBuilder.WriteString(m.theme.CorrectChar.Render(string(char)))
			} else {
				// Incorrect character - show the expected char in error style
				lineBuilder.WriteString(m.theme.IncorrectChar.Render(string(char)))
			}
		} else if i == len(m.userInput) {
			// Current cursor position
			lineBuilder.WriteString(m.theme.Cursor.Render(string(char)))
		} else {
			// Remaining characters
			lineBuilder.WriteString(m.theme.RemainingChar.Render(string(char)))
		}
	}

	// Show extra characters if user typed too much
	if len(m.userInput) > len(currentCode) {
		extra := m.userInput[len(currentCode):]
		lineBuilder.WriteString(m.theme.ExtraChar.Render(extra))
	}

	// Show cursor if at end of line
	if len(m.userInput) == len(currentCode) {
		lineBuilder.WriteString(m.theme.Cursor.Render("█"))
	}

	// Apply current line highlighting to the entire line
	return m.theme.CurrentLine.Render(lineBuilder.String())
}

// renderCompletedLineWithColors renders a completed line with color feedback
func (m *Model) renderCompletedLineWithColors(lineNum, originalCode, userInput string) string {
	// We need to get the trimmed version of the original code for comparison
	trimmedCode := strings.TrimLeft(originalCode, " \t")
	leadingSpaces := len(originalCode) - len(trimmedCode)

	// Build the display line: line number + separator + styled content
	var lineBuilder strings.Builder
	lineBuilder.WriteString(lineNum)
	lineBuilder.WriteString(" │ ")

	// Add the leading spaces (show them as dim/gray)
	if leadingSpaces > 0 {
		indentation := originalCode[:leadingSpaces]
		lineBuilder.WriteString(m.theme.RemainingChar.Render(indentation))
	}

	// Compare user input with the trimmed code and style accordingly
	maxLen := len(trimmedCode)
	if len(userInput) > maxLen {
		maxLen = len(userInput)
	}

	for i := 0; i < maxLen; i++ {
		if i < len(userInput) && i < len(trimmedCode) {
			// Character exists in both user input and expected code
			userChar := rune(userInput[i])
			expectedChar := rune(trimmedCode[i])

			if userChar == expectedChar {
				// Correct character - green
				lineBuilder.WriteString(m.theme.CorrectChar.Render(string(expectedChar)))
			} else {
				// Incorrect character - red (show the expected character)
				lineBuilder.WriteString(m.theme.IncorrectChar.Render(string(expectedChar)))
			}
		} else if i < len(trimmedCode) {
			// User didn't type this character - show it as missing/gray
			lineBuilder.WriteString(m.theme.RemainingChar.Render(string(trimmedCode[i])))
		} else {
			// User typed extra characters - show them as extra/error
			lineBuilder.WriteString(m.theme.ExtraChar.Render(string(userInput[i])))
		}
	}

	// Apply normal code line styling (no current line highlighting)
	return m.theme.CodeLine.Render(lineBuilder.String())
}

// renderMusclePowerIndicator renders the muscle power indicator panel
func (m *Model) renderMusclePowerIndicator() string {
	powerLevel := m.mpi.GetCurrentPowerLevel()
	stats := m.mpi.GetStats()

	// Create power bar
	powerBar := m.mpi.GetPowerBar(20)

	// Build MPI display
	mpiInfo := []string{
		fmt.Sprintf("%s %s", powerLevel.Icon, powerLevel.Message),
		fmt.Sprintf("💪 Power: [%s] %.0f%%", powerBar, powerLevel.Percentage),
		fmt.Sprintf("🔥 Streak: %d", stats["correct_streak"]),
		fmt.Sprintf("⚡ Peak: %.0f", stats["peak_power"]),
	}

	// Special achievement messages
	if streak, ok := stats["correct_streak"].(int); ok && streak > 0 {
		if streak >= 100 {
			mpiInfo = append(mpiInfo, "🏆 FINGER FURY UNLEASHED!")
		} else if streak >= 50 {
			mpiInfo = append(mpiInfo, "🔥 ON FIRE!")
		} else if streak >= 25 {
			mpiInfo = append(mpiInfo, "⚡ GAINING MOMENTUM!")
		}
	}

	// Fatigue warning
	if fatigueDetected, ok := stats["fatigue_detected"].(bool); ok && fatigueDetected {
		mpiInfo = append(mpiInfo, "💤 Consider a short break!")
	}

	content := strings.Join(mpiInfo, " │ ")

	// Style based on power level
	var styledContent string
	switch powerLevel.Status {
	case 0: // ZenMode
		styledContent = m.theme.Text.Foreground(lipgloss.Color("13")).Render(content) // Purple
	case 1: // FullPower
		styledContent = m.theme.Text.Foreground(lipgloss.Color("10")).Render(content) // Green
	case 2: // GoodFlow
		styledContent = m.theme.Text.Foreground(lipgloss.Color("12")).Render(content) // Blue
	case 3: // Fatigue
		styledContent = m.theme.Text.Foreground(lipgloss.Color("11")).Render(content) // Yellow
	case 4: // Burnout
		styledContent = m.theme.Text.Foreground(lipgloss.Color("9")).Render(content) // Red
	default:
		styledContent = m.theme.Text.Render(content)
	}

	return m.theme.MetricsPanel.Width(m.width - 2).Render(styledContent)
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

	// Get MPI final stats
	mpiStats := m.mpi.GetStats()
	finalPowerLevel := m.mpi.GetCurrentPowerLevel()

	stats := []string{
		fmt.Sprintf("📁 File: %s", m.filename),
		fmt.Sprintf("📄 Lines completed: %d", m.totalLines),
		fmt.Sprintf("⏱️  Total time: %s", formatDuration(m.finalStats.TotalTime)),
		fmt.Sprintf("🎯 Final accuracy: %.1f%%", m.finalStats.Accuracy),
		fmt.Sprintf("⚡ Average WPM: %.1f", m.finalStats.WPM),
		fmt.Sprintf("📊 Average CPM: %.1f", m.finalStats.CPM),
		fmt.Sprintf("❌ Total mistakes: %d", m.finalStats.TotalMistakes),
		"",
		"💪 MUSCLE POWER INDICATOR RESULTS:",
		fmt.Sprintf("🏆 Final Power State: %s %s", finalPowerLevel.Icon, finalPowerLevel.Message),
		fmt.Sprintf("⚡ Peak Power: %.1f CPM", mpiStats["peak_power"]),
		fmt.Sprintf("🔥 Max Streak: %d characters", mpiStats["max_streak"]),
		fmt.Sprintf("📈 Final Stamina: %.1f%%", mpiStats["stamina_level"].(float64)*100),
		fmt.Sprintf("🎯 Consistency Score: %.1f%%", mpiStats["consistency_score"].(float64)*100),
		fmt.Sprintf("⌨️  Total Keystrokes: %d", mpiStats["total_keystrokes"]),
		fmt.Sprintf("⏱️  Avg Keystroke Delay: %dms", mpiStats["avg_keystroke_delay"]),
	}

	// Add special achievements
	if maxStreak, ok := mpiStats["max_streak"].(int); ok && maxStreak > 0 {
		if maxStreak >= 100 {
			stats = append(stats, "🏆 ACHIEVEMENT: FINGER FURY UNLEASHED!")
		} else if maxStreak >= 50 {
			stats = append(stats, "🔥 ACHIEVEMENT: ON FIRE!")
		} else if maxStreak >= 25 {
			stats = append(stats, "⚡ ACHIEVEMENT: GAINING MOMENTUM!")
		}
	}

	// Add fatigue analysis
	if fatigueDetected, ok := mpiStats["fatigue_detected"].(bool); ok {
		if fatigueDetected {
			stats = append(stats, "💤 Fatigue was detected during session - consider breaks!")
		} else {
			stats = append(stats, "💪 Great endurance - no fatigue detected!")
		}
	}

	stats = append(stats, "", "🏆 Great job! Keep practicing to improve your speed and accuracy.")

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
