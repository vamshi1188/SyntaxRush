package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"syntaxrush/ui"
)

// Flag variables
var (
	quick      bool
	mute       bool
	stats      bool
	difficulty string
)

var practiceCmd = &cobra.Command{
	Use:   "practice [file]",
	Short: "Start typing practice session",
	Long: `Start a typing practice session with the specified file.
You can provide a file path or use shortcuts:
  go, python, js, cpp - Use built-in sample files
  
Examples:
  syntaxrush practice                    # Use default sample
  syntaxrush practice main.go            # Practice with main.go (from current directory)
  syntaxrush practice /path/to/file.go   # Practice with absolute path
  syntaxrush practice go                 # Use Go sample
  syntaxrush practice python --quick     # Quick Python practice`,
	Args: cobra.MaximumNArgs(1),
	Run:  runPractice,
}

func init() {
	rootCmd.AddCommand(practiceCmd)

	// Add flags
	practiceCmd.Flags().BoolVarP(&quick, "quick", "q", false, "Skip welcome screen and start immediately")
	practiceCmd.Flags().BoolVarP(&mute, "mute", "m", false, "Disable audio feedback")
	practiceCmd.Flags().BoolVarP(&stats, "stats", "s", false, "Show detailed stats after session")
	practiceCmd.Flags().StringVarP(&difficulty, "difficulty", "d", "normal", "Set difficulty level (easy, normal, hard)")
}

func runPractice(cmd *cobra.Command, args []string) {
	// Create a new model
	model := ui.NewModel()

	// Apply CLI flags
	if mute {
		model.SetAudioEnabled(false)
	}

	// Handle file argument
	var filePath string
	if len(args) > 0 {
		filePath = args[0]
	} else {
		filePath = "go" // Default to Go sample
	}

	// Expand the file path (handles shortcuts and resolves paths)
	resolvedPath := expandFilePath(filePath)

	// Try to load the file
	if resolvedPath != "" {
		err := model.LoadFile(resolvedPath)
		if err != nil {
			displayBanner()
			fmt.Printf("❌ Error loading file '%s': %v\n", filePath, err)
			fmt.Printf("💡 Make sure the file exists and is readable\n")
			fmt.Printf("   Current directory: %s\n", getCurrentDir())
			os.Exit(1)
		}
	}

	// Start directly if quick flag is set
	if quick {
		model.StartPracticeDirectly()
	}

	// Display banner unless in quick mode
	if !quick {
		displayBanner()
	}

	// Start the Bubble Tea program
	p := tea.NewProgram(model, tea.WithAltScreen())
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}

	// Display final stats if enabled
	if stats {
		displayFinalStats(finalModel.(*ui.Model))
	}

	// Cleanup
	finalModel.(*ui.Model).Cleanup()
}

// expandFilePath expands shortcuts and handles relative/absolute paths
func expandFilePath(input string) string {
	// Handle built-in sample shortcuts first
	switch strings.ToLower(input) {
	case "go", "sample.go":
		return findAssetFile("sample.go")
	case "py", "python", "sample.py":
		return findAssetFile("sample.py")
	case "js", "javascript", "sample.js":
		return findAssetFile("sample.js")
	case "cpp", "c++", "sample.cpp":
		return findAssetFile("sample.cpp")
	default:
		// For regular file paths, check if file exists
		if filepath.IsAbs(input) {
			// Absolute path - use as-is
			return input
		}

		// Relative path - resolve from current directory
		if absPath, err := filepath.Abs(input); err == nil {
			return absPath
		}

		// If all else fails, return input as-is
		return input
	}
}

// findAssetFile finds the asset file from the binary's location or installation
func findAssetFile(filename string) string {
	// Get the directory where the binary is located
	execPath, err := os.Executable()
	if err != nil {
		// Fallback to relative path
		return filepath.Join("assets", filename)
	}

	execDir := filepath.Dir(execPath)

	// Try different possible locations for assets
	possiblePaths := []string{
		// Assets in same directory as binary (for development)
		filepath.Join(execDir, "assets", filename),
		// Assets relative to binary (for installed version)
		filepath.Join(execDir, "..", "share", "syntaxrush", "assets", filename),
		// System-wide assets
		filepath.Join("/usr", "share", "syntaxrush", "assets", filename),
		// Local share
		filepath.Join("/usr", "local", "share", "syntaxrush", "assets", filename),
		// Fallback to current directory
		filepath.Join("assets", filename),
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// If no asset file found, return the relative path as fallback
	return filepath.Join("assets", filename)
}

// displayBanner shows the SyntaxRush banner
func displayBanner() {
	fmt.Println("🚀 SyntaxRush - Elite Code Typing Trainer")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

// getCurrentDir returns the current working directory
func getCurrentDir() string {
	if dir, err := os.Getwd(); err == nil {
		return dir
	}
	return "unknown"
}

// displayFinalStats shows final session statistics
func displayFinalStats(model *ui.Model) {
	stats := model.GetFinalStats()
	mpiStats := model.GetMPIStats()

	fmt.Println("\n🏆 Session Complete!")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("⚡ WPM: %.1f\n", stats.WPM)
	fmt.Printf("📊 CPM: %.1f\n", stats.CPM)
	fmt.Printf("🎯 Accuracy: %.1f%%\n", stats.Accuracy)
	fmt.Printf("⏱️  Duration: %v\n", stats.TotalTime)
	fmt.Printf("❌ Mistakes: %d\n", stats.TotalMistakes)

	if power, ok := mpiStats["current_power"].(float64); ok {
		fmt.Printf("💪 Final Power: %.0f%%\n", power)
	}
}
