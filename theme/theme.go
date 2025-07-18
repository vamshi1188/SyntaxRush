package theme

import "github.com/charmbracelet/lipgloss"

// Theme contains all styling for the application
type Theme struct {
	// General styles
	Title    lipgloss.Style
	Text     lipgloss.Style
	Error    lipgloss.Style
	Header   lipgloss.Style
	Controls lipgloss.Style

	// Code display styles
	CodePane    lipgloss.Style
	CodeLine    lipgloss.Style
	CurrentLine lipgloss.Style
	PaneTitle   lipgloss.Style

	// Input styles
	InputPane     lipgloss.Style
	CorrectChar   lipgloss.Style
	IncorrectChar lipgloss.Style
	RemainingChar lipgloss.Style
	ExtraChar     lipgloss.Style
	Cursor        lipgloss.Style

	// Metrics styles
	MetricsPanel lipgloss.Style
	Summary      lipgloss.Style
}

// NewDarkTheme creates a dark theme with enhanced color scheme
func NewDarkTheme() *Theme {
	return &Theme{
		// General styles
		Title: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Bold(true).
			Padding(1, 2),

		Text: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAAAAA")),

		Error: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5555")).
			Bold(true),

		Header: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00BFFF")).
			Background(lipgloss.Color("#1e1e1e")).
			Padding(0, 1).
			Bold(true),

		Controls: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAAAAA")).
			Italic(true),

		// Code display styles
		CodePane: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#555555")).
			Padding(1),

		CodeLine: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAAAAA")),

		CurrentLine: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Background(lipgloss.Color("#2a2a2a")).
			Bold(true),

		PaneTitle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00BFFF")).
			Bold(true).
			Padding(0, 1),

		// Input styles
		InputPane: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#555555")).
			Padding(1),

		CorrectChar: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF88")).
			Background(lipgloss.Color("#1e1e1e")),

		IncorrectChar: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5555")).
			Background(lipgloss.Color("#3a1a1a")),

		RemainingChar: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")),

		ExtraChar: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5555")).
			Background(lipgloss.Color("#3a1a1a")).
			Underline(true),

		Cursor: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Background(lipgloss.Color("#555555")).
			Blink(true),

		// Metrics styles
		MetricsPanel: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00BFFF")).
			Background(lipgloss.Color("#1e1e1e")).
			Padding(0, 1).
			Bold(true),

		Summary: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF88")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00FF88")).
			Padding(1, 2),
	}
}

// NewLightTheme creates a light theme
func NewLightTheme() *Theme {
	return &Theme{
		// General styles
		Title: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7C3AED")).
			Bold(true).
			Padding(1, 2),

		Text: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#374151")),

		Error: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#DC2626")).
			Bold(true),

		Header: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#1D4ED8")).
			Background(lipgloss.Color("#F1F5F9")).
			Padding(0, 1).
			Bold(true),

		Controls: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B7280")).
			Italic(true),

		// Code display styles
		CodePane: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#D1D5DB")).
			Padding(1),

		CodeLine: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#374151")),

		CurrentLine: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#92400E")).
			Background(lipgloss.Color("#FEF3C7")).
			Bold(true),

		PaneTitle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#059669")).
			Bold(true).
			Padding(0, 1),

		// Input styles
		InputPane: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#D1D5DB")).
			Padding(1),

		CorrectChar: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#065F46")).
			Background(lipgloss.Color("#D1FAE5")),

		IncorrectChar: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#991B1B")).
			Background(lipgloss.Color("#FEE2E2")),

		RemainingChar: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B7280")),

		ExtraChar: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#991B1B")).
			Background(lipgloss.Color("#FEE2E2")).
			Underline(true),

		Cursor: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#92400E")).
			Background(lipgloss.Color("#FCD34D")).
			Blink(true),

		// Metrics styles
		MetricsPanel: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#1D4ED8")).
			Background(lipgloss.Color("#F1F5F9")).
			Padding(0, 1).
			Bold(true),

		Summary: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#059669")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#10B981")).
			Padding(1, 2),
	}
}
