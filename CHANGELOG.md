# Changelog

All notable changes to SyntaxRush will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-07-20

### Added
- **Global CLI Tool**: Complete transformation from TUI app to professional CLI tool using Cobra framework
- **Global Installation**: Install SyntaxRush globally with `syntaxrush` command available from anywhere
- **Advanced CLI Commands**:
  - `syntaxrush practice [file]` - Start typing practice with built-in samples or custom files
  - `syntaxrush stats` - View comprehensive performance statistics
  - `syntaxrush config` - Configure application settings
  - `syntaxrush version` - Show detailed version information
  - `syntaxrush completion` - Generate shell auto-completion scripts
- **Professional CLI Flags**:
  - `--quick` - Skip welcome screen and start typing immediately
  - `--mute` - Disable audio feedback
  - `--stats` - Show detailed session statistics
  - `--difficulty` - Set difficulty level
- **Smart File Handling**:
  - Shortcuts for built-in samples (go, python, js, cpp)
  - Absolute and relative path support
  - Intelligent asset file location discovery
- **Shell Integration**:
  - Auto-completion support for bash, zsh, and fish
  - Tab completion for commands, files, and flags
- **Installation Scripts**:
  - Interactive installer with multiple options
  - Global installation script with asset management
  - Shell completion setup script
- **Comprehensive Documentation**:
  - Installation guide
  - Usage documentation
  - Contributing guidelines
  - Clean project structure

### Enhanced
- **Muscle Power Indicator (MPI)**: Real-time typing power tracking with 6 dynamic states
- **Audio System**: Premium 44.1kHz audio feedback with graceful fallback
- **Achievement System**: Finger Fury, On Fire, and Zen Mode achievements
- **Multi-language Support**: Built-in samples for Go, Python, JavaScript, and C++
- **Performance Metrics**: Real-time WPM, CPM, accuracy tracking
- **Error Handling**: Robust error handling with user-friendly messages
- **File Loading**: Support for any text file with UTF-8 encoding

### Technical Improvements
- **Modern CLI Architecture**: Built with Cobra framework for professional command structure
- **Modular Design**: Clean separation of concerns with cmd/, core/, ui/, theme/ packages
- **Cross-platform Support**: Works on Linux, macOS, and Windows
- **Efficient Building**: Optimized build process with stripped binaries
- **Clean Codebase**: Comprehensive documentation and code organization

### Project Structure
- Organized into logical directories: cmd/, core/, ui/, theme/, assets/, scripts/, docs/
- Professional installation and setup scripts
- Comprehensive documentation suite
- Clean .gitignore for development workflow

## [0.x.x] - Previous Versions

### Legacy Features (Pre-1.0.0)
- Basic Terminal User Interface (TUI) application
- Simple file loading and typing practice
- Basic metrics tracking
- Original Muscle Power Indicator concept
- Audio feedback system foundation
- Core typing practice functionality

---

## Version Schema

- **Major versions** (x.0.0): Breaking changes, major new features
- **Minor versions** (1.x.0): New features, enhancements, non-breaking changes  
- **Patch versions** (1.0.x): Bug fixes, minor improvements

## Development

See [CONTRIBUTING.md](docs/CONTRIBUTING.md) for development guidelines and contribution instructions.
