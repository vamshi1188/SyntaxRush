package main

import (
	"fmt"
	"log"
	"os"

	"syntaxrush/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Initialize the app
	model := ui.NewModel()

	// Check if file was provided as argument
	if len(os.Args) > 1 {
		err := model.LoadFile(os.Args[1])
		if err != nil {
			fmt.Printf("Error loading file: %v\n", err)
			os.Exit(1)
		}
	}

	// Create the Bubble Tea program
	p := tea.NewProgram(model, tea.WithAltScreen())

	// Run the program
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
