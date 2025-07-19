# Usage Guide

## Basic Commands

### Practice Command

The core command for typing practice:

```bash
# Use built-in samples
syntaxrush practice go          # Go sample
syntaxrush practice python      # Python sample
syntaxrush practice js          # JavaScript sample
syntaxrush practice cpp         # C++ sample

# Practice with your own files
syntaxrush practice main.go     # Relative path
syntaxrush practice /path/to/file.py  # Absolute path

# Quick start (skip welcome screen)
syntaxrush practice go --quick

# Silent mode (no audio)
syntaxrush practice --mute

# Show detailed stats after session
syntaxrush practice --stats
```

### Other Commands

```bash
# View performance statistics
syntaxrush stats

# Configure settings
syntaxrush config

# Version information
syntaxrush version

# Generate shell completion
syntaxrush completion bash > completion.sh
```

## Features

### Muscle Power Indicator (MPI)
- Real-time typing power tracking
- 6 dynamic states: Ready, Warming Up, On Fire, Finger Fury, Fatigue, Recovery
- Visual power bar shows your current typing state

### Audio Feedback
- High-quality error beeps for mistakes
- Success sounds for completed lines
- Can be disabled with `--mute` flag

### Achievement System
- **Finger Fury**: Sustained high-speed typing
- **On Fire**: Extended accurate typing streaks
- **Zen Mode**: Perfect accuracy sessions

### Multi-language Support
Built-in samples for:
- Go (.go)
- Python (.py)
- JavaScript (.js)
- C++ (.cpp)

### Real-time Metrics
- **WPM**: Words per minute
- **CPM**: Characters per minute
- **Accuracy**: Percentage of correct keystrokes
- **Mistakes**: Error count and tracking
- **Time**: Session duration

## Keyboard Shortcuts

During typing practice:
- `Enter`: Complete current line
- `Ctrl+R`: Retry/restart session
- `Ctrl+U`: Upload new file
- `Esc`: Return to menu
- `Ctrl+C`: Quit application

Note: Backspace is disabled to encourage accuracy!

## Tips for Better Performance

1. **Start Slow**: Focus on accuracy first, speed comes naturally
2. **Use Proper Posture**: Sit up straight, hands positioned correctly
3. **Regular Practice**: Short, frequent sessions are more effective
4. **Learn from Mistakes**: Pay attention to error patterns
5. **Use MPI**: Watch your power indicator to maintain optimal typing state

## File Formats

SyntaxRush works with any text file, but is optimized for:
- Source code files (.go, .py, .js, .cpp, .rs, .java)
- Plain text files
- Configuration files
- Documentation files

Files should be UTF-8 encoded for best results.
