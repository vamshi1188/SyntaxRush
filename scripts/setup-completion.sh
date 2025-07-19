#!/bin/bash

# SyntaxRush Shell Completion Setup Script
# Enables auto-completion for bash, zsh, and fish shells

set -e

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

# Detect shell
detect_shell() {
    if [[ -n "$ZSH_VERSION" ]]; then
        CURRENT_SHELL="zsh"
    elif [[ -n "$BASH_VERSION" ]]; then
        CURRENT_SHELL="bash"
    else
        CURRENT_SHELL=$(basename "$SHELL")
    fi
    
    print_status "Detected shell: $CURRENT_SHELL"
}

# Setup bash completion
setup_bash() {
    print_status "Setting up bash completion..."
    
    # Check if syntaxrush is available
    if ! command -v syntaxrush &> /dev/null; then
        print_warning "syntaxrush command not found. Please install it first with ./install.sh"
        return 1
    fi
    
    # Generate completion script
    syntaxrush completion bash > syntaxrush_completion.bash
    
    # Try to install system-wide
    if [[ -d "/usr/share/bash-completion/completions" ]]; then
        sudo mv syntaxrush_completion.bash /usr/share/bash-completion/completions/syntaxrush
        print_success "Bash completion installed system-wide"
    elif [[ -d "/etc/bash_completion.d" ]]; then
        sudo mv syntaxrush_completion.bash /etc/bash_completion.d/syntaxrush
        print_success "Bash completion installed system-wide"
    else
        # Install to user directory
        mkdir -p ~/.local/share/bash-completion/completions
        mv syntaxrush_completion.bash ~/.local/share/bash-completion/completions/syntaxrush
        
        # Add to bashrc if not already there
        if ! grep -q "bash-completion" ~/.bashrc; then
            echo "" >> ~/.bashrc
            echo "# Enable bash completion" >> ~/.bashrc
            echo "[[ -r ~/.local/share/bash-completion/completions ]] && source ~/.local/share/bash-completion/completions/*" >> ~/.bashrc
        fi
        
        print_success "Bash completion installed to user directory"
        print_warning "Please restart your shell or run: source ~/.bashrc"
    fi
}

# Setup zsh completion
setup_zsh() {
    print_status "Setting up zsh completion..."
    
    if ! command -v syntaxrush &> /dev/null; then
        print_warning "syntaxrush command not found. Please install it first with ./install.sh"
        return 1
    fi
    
    # Create completion directory
    mkdir -p ~/.config/zsh/completions
    
    # Generate completion script
    syntaxrush completion zsh > ~/.config/zsh/completions/_syntaxrush
    
    # Add to zshrc if not already there
    if ! grep -q "~/.config/zsh/completions" ~/.zshrc; then
        echo "" >> ~/.zshrc
        echo "# Enable zsh completion for syntaxrush" >> ~/.zshrc
        echo "fpath=(~/.config/zsh/completions \$fpath)" >> ~/.zshrc
        echo "autoload -U compinit && compinit" >> ~/.zshrc
    fi
    
    print_success "Zsh completion installed"
    print_warning "Please restart your shell or run: source ~/.zshrc"
}

# Setup fish completion
setup_fish() {
    print_status "Setting up fish completion..."
    
    if ! command -v syntaxrush &> /dev/null; then
        print_warning "syntaxrush command not found. Please install it first with ./install.sh"
        return 1
    fi
    
    # Create completion directory
    mkdir -p ~/.config/fish/completions
    
    # Generate completion script
    syntaxrush completion fish > ~/.config/fish/completions/syntaxrush.fish
    
    print_success "Fish completion installed"
    print_warning "Please restart your shell"
}

# Main function
main() {
    echo "🚀 SyntaxRush Shell Completion Setup"
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo ""
    
    detect_shell
    
    case "$CURRENT_SHELL" in
        bash)
            setup_bash
            ;;
        zsh)
            setup_zsh
            ;;
        fish)
            setup_fish
            ;;
        *)
            print_warning "Unsupported shell: $CURRENT_SHELL"
            echo "Supported shells: bash, zsh, fish"
            echo ""
            echo "Manual setup:"
            echo "syntaxrush completion bash > syntaxrush_completion.bash"
            echo "syntaxrush completion zsh > syntaxrush_completion.zsh"
            echo "syntaxrush completion fish > syntaxrush_completion.fish"
            exit 1
            ;;
    esac
    
    echo ""
    echo "🎉 Shell completion setup complete!"
    echo ""
    echo "Try typing 'syntaxrush ' and press Tab for auto-completion!"
    echo ""
    echo "Available completions:"
    echo "• Commands: practice, stats, config, version"
    echo "• Files: Auto-complete file paths"
    echo "• Flags: --help, --mute, --quick, --stats"
    echo "• Samples: go, python, js, cpp"
}

main "$@"
