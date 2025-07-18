# SyntaxRush - Terminal-Based Code Typing Practice Tool

🧠⚡ **SyntaxRush** is a rich, interactive terminal application built in Go that helps developers improve their coding speed and accuracy through focused typing practice with real code files.

*Practice typing real code. Master syntax.*

## Features

### 🎯 Core Functionality
- **File Upload Support**: Load code files in multiple languages (.go, .py, .js, .cpp, .java, .rs, .ts)
- **Smart File Loading**: Use shortcuts ('go', 'py', 'js', 'cpp') or full/relative paths
- **Real-time Input**: Interactive file path input with live cursor and error feedback
- **Real-time Feedback**: Instant visual feedback with color-coded correct/incorrect characters
- **Live Metrics**: Real-time WPM, CPM, accuracy, and mistake tracking
- **Syntax Highlighting**: Clean code display with current line highlighting
- **Session Statistics**: Comprehensive performance analysis after completion

### 🎨 UI Features
- **Beautiful ASCII Art**: Stunning welcome screen with SyntaxRush branding
- **Enhanced Color Scheme**: Professional dark theme with neon green/red feedback
- **Split-pane Interface**: Code display and typing input in separate panes
- **Progress Tracking**: Visual progress indicator and line-by-line navigation
- **Responsive Design**: Adapts to terminal size with smart scrolling
- **Eye-friendly Theme**: Optimized colors for extended coding sessions

### ⌨️ Controls
- **Ctrl+U**: Upload new file
- **Ctrl+R**: Retry current session
- **Enter**: Complete current line
- **Esc**: Return to menu
- **Q**: Quit application

### 📊 Metrics Tracked
- **WPM** (Words Per Minute)
- **CPM** (Characters Per Minute)  
- **Accuracy** percentage
- **Mistake count** with position tracking
- **Session time** with live timer
- **Error heatmap** for performance analysis

## Installation

### Prerequisites
- Go 1.21 or higher
- Terminal with color support

### Quick Start

1. **Clone or download the project**:
   ```bash
   # If using git
   git clone <repository-url>
   cd typingapp
   
   # Or create the project from provided files
   mkdir codetype && cd codetype
   # Copy all provided files to this directory
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the application**:
   ```bash
   # Run directly
   go run main.go
   
   # Or build and run
   go build -o syntaxrush
   ./syntaxrush
   
   # Or use the showcase launcher
   ./launch.sh
   ```

4. **Run with a specific file**:
   ```bash
   go run main.go path/to/your/code/file.go
   # or
   ./syntaxrush path/to/your/code/file.py
   ```

## Usage

### Getting Started

1. **Launch the app**:
   ```bash
   go run main.go
   ```

2. **Welcome Screen**: 
   - Press `Enter` or `Space` to start with the built-in sample
   - Press `Ctrl+U` to upload your own code file
   - Press `Q` or `Esc` to quit

3. **Typing Practice**:
   - Type each line exactly as shown in the code pane
   - Watch real-time feedback with color coding:
     - 🟢 **Green**: Correct characters
     - 🔴 **Red**: Incorrect characters  
     - ⚪ **Gray**: Characters yet to type
     - 🟡 **Yellow**: Current cursor position
   - Press `Enter` to move to the next line

4. **Session Complete**:
   - View comprehensive statistics
   - Choose to retry, upload new file, or return to menu

### Sample Files

The application includes sample files in the `assets/` directory:

- **`sample.go`**: Go calculator with structs, methods, and error handling
- **`sample.py`**: Python data processor with classes, type hints, and statistics
- **`sample.js`**: JavaScript task manager with modern ES6+ features

### Supported File Types

- **Go**: `.go`
- **Python**: `.py`
- **JavaScript**: `.js`, `.jsx`, `.ts`, `.tsx`
- **C++**: `.cpp`, `.c`
- **Java**: `.java`
- **Rust**: `.rs`

## Project Structure

```
syntaxrush/
├── main.go              # Application entry point
├── go.mod              # Go module definition  
├── build.sh            # Build script
├── launch.sh           # Demo launcher with showcase
├── demo.sh             # File upload demo
├── showcase.sh         # Color theme showcase
├── ui/
│   ├── model.go        # Bubbletea model & state management
│   └── view.go         # UI rendering with ASCII art
├── core/
│   ├── parser.go       # File parsing & validation
│   ├── metrics.go      # Statistics calculation
│   └── timer.go        # Session timing
├── theme/
│   └── theme.go        # Enhanced color theming
├── assets/
│   ├── sample.go       # Sample Go code
│   ├── sample.py       # Sample Python code
│   ├── sample.js       # Sample JavaScript code
│   └── sample.cpp      # Sample C++ code
└── README.md           # This file
```

## Technical Details

### Architecture
- **Framework**: [Bubble Tea](https://github.com/charmbracelet/bubbletea) (Go TUI framework)
- **Styling**: [Lip Gloss](https://github.com/charmbracelet/lipgloss) (Terminal styling)
- **Terminal**: [tcell](https://github.com/gdamore/tcell) (Terminal handling)

### Key Components

1. **Parser**: Handles file reading, validation, and content processing
2. **Metrics Engine**: Real-time calculation of typing statistics
3. **Timer**: Precise session timing and elapsed time tracking
4. **UI Model**: State management using the Model-View-Update pattern
5. **Theme System**: Consistent styling across all UI components

### Performance Features
- **Efficient Rendering**: Only updates changed screen regions
- **Smart Scrolling**: Keeps current line visible in large files
- **Real-time Updates**: Sub-second metric calculation and display
- **Memory Efficient**: Minimal memory footprint even with large files

## Examples

### Basic Session Flow

```
1. Launch: go run main.go
2. ASCII art welcome screen appears
3. Press Enter to start with sample.go
4. Type each line character by character
5. Watch metrics update in real-time
6. Complete session and view statistics
7. Choose next action (retry, new file, quit)
```

### Loading Custom File

```bash
# Start with specific file
go run main.go mycode.py

# Or upload during session
# Press Ctrl+U and enter file path or shortcut:
#   'go' → loads sample.go
#   'py' → loads sample.py  
#   'js' → loads sample.js
#   'cpp' → loads sample.cpp
#   './myfile.go' → loads relative path
#   '/full/path/file.py' → loads absolute path
```

### Sample Output

```
  ____                  _                  ____           _     
 / ___|  ___ _ __  _ __(_)_ __   __ _     |  _ \ ___  ___| |_  
 \___ \ / _ \ '_ \| '__| | '_ \ / _` |____| |_) / _ \/ __| __| 
  ___) |  __/ | | | |  | | | | | (_| |____|  _ <  __/\__ \ |_  
 |____/ \___|_| |_|_|  |_|_| |_|\__, |    |_| \_\___||___/\__| 
                               |___/                           

           Practice typing real code. Master syntax.

📁 sample.go • Progress: 15/45 (33.3%)

📖 Code to Type
  13 │ func NewCalculator() *Calculator {
  14 │     return &Calculator{memory: 0}
  15 │ }                                    ← Current line
  16 │ 
  17 │ // Add performs addition

⌨️  Your Input
return &Calculator{memory: 0█

⏱️ Time: 02:34 │ 🎯 Accuracy: 94.2% │ ⚡ WPM: 47 │ 📊 CPM: 235 │ ❌ Mistakes: 3

Ctrl+R: Retry │ Ctrl+U: Upload │ Esc: Menu
```

## Contributing

This is a complete, functional typing practice tool. Potential enhancements:

- **Multiple themes** (light/dark mode toggle)
- **Difficulty levels** (beginner, intermediate, advanced)
- **Leaderboards** (local score storage)
- **Custom snippets** (user-defined practice content)
- **Language-specific practice** (focus on specific syntax patterns)

## Dependencies

- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)**: Modern TUI framework
- **[Lip Gloss](https://github.com/charmbracelet/lipgloss)**: Terminal styling library
- **[tcell](https://github.com/gdamore/tcell)**: Terminal cell manipulation

## License

Open source - feel free to use, modify, and distribute.

---

**Happy Typing! 🧠⚡**

*Practice typing real code. Master syntax with SyntaxRush.*
