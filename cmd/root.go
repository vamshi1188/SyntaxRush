package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "syntaxrush",
	Short: "🚀 Elite Code Typing Trainer - Practice typing real code",
	Long: `🚀 SyntaxRush - Elite Code Typing Trainer

The ultimate muscle-powered code typing experience. Practice typing real code,
build typing stamina, and master syntax with our revolutionary Muscle Power
Indicator (MPI) system.

Features:
• 💪 Real-time Power Tracking with 6 dynamic states
• 🏆 Achievement System (Finger Fury, On Fire, Zen Mode)
• 🔊 Premium Audio Feedback (44.1kHz audio)
• 📊 Comprehensive Analytics and Health Insights
• 🎨 Live Color Feedback with typing history
• 📁 Multi-language Support (Go, Python, JS, C++, Rust, Java)

Examples:
  syntaxrush practice filename     # Practice with your code file
  syntaxrush practice go          # Practice with Go sample
  syntaxrush practice python      # Practice with Python sample
  syntaxrush stats                # View your performance stats
  syntaxrush config               # Configure settings`,
	Version: "1.0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = false
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "help [command]",
		Short:  "Help about any command",
		Hidden: true,
	})
}
