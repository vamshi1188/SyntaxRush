package core

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Parser handles file parsing and content processing
type Parser struct {
	supportedExtensions map[string]bool
}

// NewParser creates a new parser instance
func NewParser() *Parser {
	return &Parser{
		supportedExtensions: map[string]bool{
			".go":   true,
			".py":   true,
			".js":   true,
			".cpp":  true,
			".c":    true,
			".java": true,
			".rs":   true,
			".ts":   true,
			".jsx":  true,
			".tsx":  true,
		},
	}
}

// ParseFile reads and processes a code file
func (p *Parser) ParseFile(filename string) (string, error) {
	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", filename)
	}

	// Check file extension
	ext := filepath.Ext(filename)
	if !p.supportedExtensions[ext] {
		return "", fmt.Errorf("unsupported file extension: %s", ext)
	}

	// Read file
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Keep original formatting including tabs and spaces
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	return strings.Join(lines, "\n"), nil
}

// IsSupported checks if a file extension is supported
func (p *Parser) IsSupported(filename string) bool {
	ext := filepath.Ext(filename)
	return p.supportedExtensions[ext]
}

// GetSupportedExtensions returns list of supported file extensions
func (p *Parser) GetSupportedExtensions() []string {
	var extensions []string
	for ext := range p.supportedExtensions {
		extensions = append(extensions, ext)
	}
	return extensions
}
