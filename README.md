# 🚀 SyntaxRush - Elite Code Typing Trainer

> **💪 The Ultimate Muscle-Powered Code Typing Experience**  
> *Practice real code. Build typing stamina. Master syntax.*

🧠⚡ **SyntaxRush** is an advanced, gamified terminal application that transforms code typing practice into an engaging fitness experience. Built for developers who want to type code like a pro!

## ✨ Key Features

### � **Muscle Power Indicator (MPI)** - *Revolutionary Typing Endurance System*
- **💪 Real-time Power Tracking**: Dynamic muscle power calculation based on keystroke efficiency
- **🏆 Achievement System**: Unlock "Finger Fury", "On Fire", and "Zen Mode" achievements  
- **📈 Stamina Monitoring**: Track typing endurance and detect fatigue patterns
- **⚡ Power States**: 6 dynamic states from "Ready to Type" to "Burnout" with visual feedback
- **🔥 Streak Tracking**: Monitor consecutive correct characters with milestone celebrations
- **🎯 Consistency Analysis**: Rhythm detection and keystroke timing optimization

### 🧠 **Smart Code Practice**
- **📁 Multi-language Support**: Go, Python, JavaScript, C++, TypeScript, Rust, Java
- **🚀 Quick File Loading**: Smart shortcuts (`go`, `py`, `js`, `cpp`) or custom file paths
- **✨ Live Color Feedback**: Green/red character highlighting with persistent history
- **📖 Unified Display**: Code context and typing practice in one seamless interface
- **🎨 Leading Space Intelligence**: Skip indentation, focus on actual code content

### 🔊 **Premium Audio Experience**
- **🎵 High-Quality Audio**: Oto v2 library for crisp 44.1kHz sound
- **❌ Error Feedback**: Instant audio cues for typing mistakes
- **✅ Success Sounds**: Satisfying completion audio rewards
- **🔇 Graceful Fallback**: Terminal bell backup if audio fails

### 📊 **Comprehensive Analytics**
- **⏱️ Real-time Metrics**: WPM, CPM, accuracy, time tracking
- **🎯 Session Analysis**: Detailed performance breakdown
- **💾 Typing History**: Color-coded progress preservation across lines
- **📈 Power Statistics**: Peak performance, stamina levels, consistency scores

## 🎮 Power States & Achievements

| Status | Icon | Trigger Condition | Description |
|--------|------|------------------|-------------|
| 🧘 **Zen Mode** | 🧘 | 95%+ consistency + 80+ CPM + 50+ streak | Ultimate flow state |
| 💪 **Full Power** | 💪 | 80%+ stamina + 60+ CPM | Peak performance |
| ⚡ **Good Flow** | ⚡ | 60%+ stamina + steady rhythm | Solid pace |
| 💤 **Fatigue Mode** | � | Declining performance | Focus needed |
| 🔥 **Rest Needed** | 🔥 | Critical performance drop | Break time |
| 🚀 **Ready to Type** | 🚀 | Initial state | Let's begin! |

### 🏆 **Achievement Unlocks**
- **🏆 FINGER FURY UNLEASHED!** - 100+ character perfect streak
- **🔥 ON FIRE!** - 50+ character perfect streak  
- **⚡ GAINING MOMENTUM!** - 25+ character perfect streak

## 🚀 Quick Start

### Prerequisites
- **Go 1.21+** (Download from [golang.org](https://golang.org))
- **Color-capable terminal** (most modern terminals)
### Installation

```bash
# Clone the repository
git clone https://github.com/vamshi1188/SyntaxRush.git
cd SyntaxRush

# Install dependencies
go mod tidy

# Build the application
go build -o syntaxrush

# Run SyntaxRush
./syntaxrush
```

### Alternative: Direct Run
```bash
# Run without building
go run main.go

# Run with a specific file
./syntaxrush path/to/your/code.go
```

## 🎯 How to Use

### 🚀 **Getting Started**

1. **Launch SyntaxRush**:
   ```bash
   ./syntaxrush
   ```

2. **Choose your practice mode**:
   - **Quick Start**: Press `Enter` or `Space` to begin with sample code
   - **Custom File**: Press `Ctrl+U` to upload your own code file

3. **File Shortcuts** (when uploading):
   - Type `go` for Go sample
   - Type `py` for Python sample  
   - Type `js` for JavaScript sample
   - Type `cpp` for C++ sample
   - Or enter any file path: `src/main.go`, `~/project/app.py`

### 💪 **Practice Session**

1. **Start Typing**: Begin typing the displayed code
2. **Watch Your Power**: Monitor your Muscle Power Indicator in real-time
3. **Build Streaks**: Aim for long correct character streaks
4. **Maintain Flow**: Keep consistent rhythm for maximum power
5. **Complete Lines**: Press `Enter` when you finish each line

### 📊 **Understanding Your Results**

#### **Real-time Display**
```
🚀 Ready to Type │ 💪 Power: [████████████████████] 100% │ 🔥 Streak: 0 │ ⚡ Peak: 1

⏱️ Time: 00:00 │ 🎯 Accuracy: 100.0% │ ⚡ WPM: 0 │ 📊 CPM: 0 │ ❌ Mistakes: 0
```

#### **Final Session Summary**
- **📁 Session Info**: File name, lines completed, total time
- **🎯 Performance**: WPM, CPM, accuracy, mistakes
- **💪 MPI Results**: Power states, peak performance, streaks
- **🏆 Achievements**: Unlocked achievements and milestones
- **💡 Health Insights**: Fatigue detection and endurance analysis

## ⌨️ **Controls & Shortcuts**

### **Main Controls**
| Key | Action | Description |
|-----|--------|-------------|
| `Enter` / `Space` | Start Practice | Begin typing session with current file |
| `Ctrl+U` | Upload File | Load a new code file for practice |
| `Ctrl+R` | Retry Session | Restart current file from beginning |
| `Esc` | Return to Menu | Go back to welcome screen |
| `Q` | Quit | Exit SyntaxRush |

### **File Loading Shortcuts**
| Shortcut | File | Language |
|----------|------|----------|
| `go` | sample.go | Go programming |
| `py` | sample.py | Python |
| `js` | sample.js | JavaScript |
| `cpp` | sample.cpp | C++ |

### **During Practice**
- **No Backspace**: Practice forward-only typing (realistic coding)
- **Audio Feedback**: Hear mistake alerts and success sounds
- **Live Color Coding**: See your progress in real-time
- **Automatic Spacing**: Skip leading indentation, focus on code

## 🎨 **Visual Experience**

### **Color Coding System**
- **🟢 Green Characters**: Correctly typed
- **🔴 Red Characters**: Mistakes (shows expected character)
- **⚪ Gray Characters**: Not yet typed or indentation
- **🟡 Yellow Cursor**: Current typing position
- **🟠 Orange**: Extra characters beyond line end

### **Power Bar Display**
```
💪 Full Power │ 💪 Power: [████████████████████] 100% │ 🔥 Streak: 25 │ ⚡ Peak: 150
```

## 🏗️ **Technical Architecture**

### **Built With**
- **🔧 Go 1.21+**: Core application language
- **🖥️ Bubble Tea**: Elegant TUI framework  
- **🎨 Lip Gloss**: Beautiful terminal styling
- **🔊 Oto v2**: High-quality cross-platform audio
- **⚡ tcell v2**: Advanced terminal capabilities

### **Project Structure**
```
SyntaxRush/
├── main.go              # Application entry point
├── ui/                  # User interface components
│   ├── model.go         # Application state & logic
│   └── view.go          # UI rendering & layout
├── core/                # Core functionality
│   ├── audio.go         # Audio management
│   ├── metrics.go       # Performance tracking
│   ├── muscle_power.go  # MPI system
│   ├── parser.go        # File parsing
│   └── timer.go         # Time management
├── theme/               # Visual theming
│   └── theme.go         # Color schemes & styles
├── assets/              # Sample files
│   ├── sample.go        # Go calculator
│   ├── sample.py        # Python data processor  
│   ├── sample.js        # JavaScript task manager
│   └── sample.cpp       # C++ grade system
└── README.md            # This file
```

## 📚 **Sample Files**

## 🚧 **Supported Languages**

SyntaxRush supports practice with these programming languages:

| Language | Extensions | Sample File |
|----------|------------|-------------|
| **Go** | `.go` | Advanced calculator with structs & methods |
| **Python** | `.py` | Data processor with classes & statistics |
| **JavaScript** | `.js`, `.jsx`, `.ts`, `.tsx` | Task manager with ES6+ features |
| **C++** | `.cpp`, `.c`, `.cc` | Grade system with OOP principles |
| **Java** | `.java` | *Add your own files* |
| **Rust** | `.rs` | *Add your own files* |

## 🔧 **Development**

### **Contributing**
1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit changes: `git commit -m 'Add amazing feature'`
4. Push to branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

### **Building from Source**
```bash
# Clone and build
git clone https://github.com/vamshi1188/SyntaxRush.git
cd SyntaxRush
go mod tidy
go build -o syntaxrush

# Run tests
go test ./...

# Run with race detection
go run -race main.go
```

## 🎯 **Why SyntaxRush?**

### **🚀 For Developers**
- **Real Code Practice**: No more lorem ipsum - practice with actual code
- **Language Agnostic**: Works with any programming language
- **Muscle Memory**: Build instinctive syntax knowledge
- **Flow State Training**: Develop sustained coding rhythm

### **🎮 For Gamers**
- **Achievement System**: Unlock progressively harder challenges
- **Power Progression**: Build typing endurance like a fitness tracker  
- **Visual Feedback**: Satisfying real-time progress indicators
- **Competitive Elements**: Beat your personal bests

### **💪 For Health**
- **Fatigue Detection**: Prevents typing strain and RSI
- **Break Reminders**: Promotes healthy practice habits
- **Stamina Building**: Gradual endurance improvement
- **Rhythm Training**: Develops consistent, sustainable pace

## 📜 **License**

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 **Acknowledgments**

- **[Charm](https://charm.sh/)** - For the amazing Bubble Tea and Lip Gloss libraries
- **[Hajimehoshi](https://github.com/hajimehoshi)** - For the Oto audio library
- **Go Community** - For excellent tooling and ecosystem
- **Contributors** - Thank you for making SyntaxRush better!

---

<div align="center">

**🚀 Ready to become a typing master?**

[Download SyntaxRush](https://github.com/vamshi1188/SyntaxRush/releases) • [Report Bug](https://github.com/vamshi1188/SyntaxRush/issues) • [Request Feature](https://github.com/vamshi1188/SyntaxRush/issues)

**Built with 💪 for developers who want to type like pros**

</div>
