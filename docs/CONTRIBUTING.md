# Contributing to SyntaxRush

Thank you for your interest in contributing to SyntaxRush! This guide will help you get started.

## Development Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/vamshi1188/SyntaxRush.git
   cd SyntaxRush
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Build and test**:
   ```bash
   go build -o syntaxrush .
   ./syntaxrush practice go
   ```

## Project Structure

```
SyntaxRush/
├── cmd/                    # CLI commands (Cobra)
│   ├── root.go            # Root command and global flags
│   ├── practice.go        # Practice command implementation
│   └── commands.go        # Other commands (stats, config, version)
├── core/                  # Core business logic
│   ├── metrics.go         # Performance tracking
│   ├── parser.go          # File parsing
│   ├── timer.go           # Time management
│   └── audio.go           # Audio feedback system
├── ui/                    # User interface (Bubble Tea)
│   ├── model.go           # Application state and logic
│   └── view.go            # Rendering and display
├── theme/                 # UI themes and styling
│   └── theme.go           # Color schemes and styles
├── assets/                # Sample files for practice
│   ├── sample.go
│   ├── sample.py
│   ├── sample.js
│   └── sample.cpp
├── scripts/               # Installation and setup scripts
│   ├── install-global.sh
│   └── setup-completion.sh
├── docs/                  # Documentation
│   ├── INSTALL.md
│   ├── USAGE.md
│   └── CONTRIBUTING.md
└── main.go               # Application entry point
```

## How to Contribute

### 1. Reporting Issues

- Use the GitHub issue tracker
- Provide clear reproduction steps
- Include system information (OS, Go version)
- Attach screenshots for UI issues

### 2. Feature Requests

- Describe the problem you're trying to solve
- Explain how the feature would help users
- Consider implementation complexity

### 3. Code Contributions

1. **Fork the repository**
2. **Create a feature branch**:
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Make your changes**
4. **Test thoroughly**:
   ```bash
   go test ./...
   go build -o syntaxrush .
   ./syntaxrush practice go
   ```
5. **Submit a pull request**

### 4. Adding New Sample Files

To add support for a new programming language:

1. **Add sample file**:
   ```bash
   # Create assets/sample.rust
   # Add well-commented, representative code
   ```

2. **Update practice command**:
   ```go
   // In cmd/practice.go, add to expandFilePath():
   case "rust", "rs", "sample.rust":
       return findAssetFile("sample.rust")
   ```

3. **Test the integration**:
   ```bash
   go build -o syntaxrush .
   ./syntaxrush practice rust
   ```

## Code Style Guidelines

### Go Code Style
- Follow standard Go formatting (`go fmt`)
- Use meaningful variable and function names
- Add comments for exported functions
- Keep functions focused and small
- Handle errors gracefully

### Git Commit Messages
- Use present tense ("Add feature" not "Added feature")
- Use imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit first line to 72 characters
- Reference issues and pull requests when applicable

### Example Commit Messages
```
Add Rust language support

- Add sample.rust file with common patterns
- Update practice command to recognize rust/rs shortcuts
- Add Rust-specific syntax highlighting

Fixes #42
```

## Testing Guidelines

### Manual Testing
- Test all CLI commands and flags
- Verify audio feedback works (and graceful fallback)
- Test with various file types and sizes
- Verify global installation works correctly

### Automated Testing
- Add unit tests for new core functionality
- Test error handling paths
- Verify metrics calculations are accurate

## Release Process

1. Update version in `cmd/root.go`
2. Update `CHANGELOG.md`
3. Create release tag
4. Build binaries for multiple platforms
5. Update installation scripts if needed

## Getting Help

- Check existing issues and documentation
- Ask questions in GitHub Discussions
- Join our community chat (if available)
- Mention @vamshi1188 in issues for urgent matters

## Recognition

Contributors will be:
- Listed in `CONTRIBUTORS.md`
- Mentioned in release notes
- Given credit in relevant documentation

Thank you for helping make SyntaxRush better! 🚀
