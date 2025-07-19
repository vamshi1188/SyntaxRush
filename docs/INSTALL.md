# Installation Guide

## Quick Install (Recommended)

```bash
# Clone the repository
git clone https://github.com/vamshi1188/SyntaxRush.git
cd SyntaxRush

# Run the interactive installer
./install.sh
```

## Installation Methods

### 1. Global Installation (Local Source)

Install SyntaxRush globally from your local source code:

```bash
./scripts/install-global.sh
```

After installation, you can use SyntaxRush from any directory:

```bash
syntaxrush practice main.go
syntaxrush practice go --quick
syntaxrush stats
```

### 2. Development Build

For development or local testing:

```bash
go build -o syntaxrush .
./syntaxrush practice go
```

### 3. Shell Completion

Enable auto-completion for enhanced CLI experience:

```bash
./scripts/setup-completion.sh
```

Supports bash, zsh, and fish shells.

### 4. GitHub Installation (Coming Soon)

Once the updated code is pushed to GitHub, users can install directly:

```bash
go install github.com/vamshi1188/SyntaxRush@latest
```

**Important**: Use the exact case-sensitive path:
- ✅ Correct: `github.com/vamshi1188/SyntaxRush`
- ❌ Incorrect: `github.com/vamshi1188/Syntaxrush`

## Requirements

- Go 1.21 or higher
- Linux/macOS/Windows
- Terminal with Unicode support

## Verification

Test your installation:

```bash
syntaxrush version
syntaxrush --help
syntaxrush practice go --quick
```

## Troubleshooting

### Command not found
If `syntaxrush` command is not found after global installation:

1. Check if `/usr/local/bin` is in your PATH:
   ```bash
   echo $PATH | grep -q "/usr/local/bin" && echo "OK" || echo "Missing"
   ```

2. Add to PATH if missing:
   ```bash
   echo 'export PATH="$PATH:/usr/local/bin"' >> ~/.bashrc
   source ~/.bashrc
   ```

### Permission denied
If you get permission errors:

```bash
# Make scripts executable
chmod +x scripts/*.sh
chmod +x install.sh
```

### Audio issues
If audio doesn't work:
- SyntaxRush gracefully falls back to visual feedback
- Check if your system supports audio output
- Use `--mute` flag to disable audio: `syntaxrush practice --mute`
