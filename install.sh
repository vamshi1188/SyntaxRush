#!/bin/bash

# SyntaxRush Quick Install Script
# This script provides easy installation options for SyntaxRush

set -e

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${BLUE}🚀 SyntaxRush Installation${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "Choose installation method:"
echo ""
echo "1. Global Installation (Recommended)"
echo "   - Install syntaxrush globally to /usr/local/bin"
echo "   - Available from anywhere: 'syntaxrush practice file.go'"
echo "   - Installs from local source code"
echo ""
echo "2. Shell Completion Setup"
echo "   - Add auto-completion for bash/zsh/fish"
echo "   - Enables tab completion for commands and files"
echo ""
echo "3. Development Build"
echo "   - Build locally for development"
echo "   - Run with './syntaxrush'"
echo ""
echo "4. GitHub Installation (Future)"
echo "   - Install directly from GitHub (after code is pushed)"
echo "   - Command: go install github.com/vamshi1188/SyntaxRush@latest"
echo ""

read -p "Enter your choice (1-4): " choice

case $choice in
    1)
        echo -e "${BLUE}🔧 Running global installation...${NC}"
        ./scripts/install-global.sh
        ;;
    2)
        echo -e "${BLUE}🔧 Setting up shell completion...${NC}"
        ./scripts/setup-completion.sh
        ;;
    3)
        echo -e "${BLUE}🔧 Building for development...${NC}"
        go build -o syntaxrush .
        echo -e "${GREEN}✅ Build complete! Run with: ./syntaxrush${NC}"
        echo -e "${YELLOW}💡 To test: ./syntaxrush practice go --quick${NC}"
        ;;
    4)
        echo -e "${YELLOW}📋 GitHub Installation Instructions:${NC}"
        echo ""
        echo "After pushing your code to GitHub, users can install with:"
        echo -e "${BLUE}go install github.com/vamshi1188/SyntaxRush@latest${NC}"
        echo ""
        echo "Note: Make sure to use the exact case-sensitive path:"
        echo "✅ Correct:   github.com/vamshi1188/SyntaxRush"
        echo "❌ Incorrect: github.com/vamshi1188/Syntaxrush"
        echo ""
        echo "Current status: Repository needs to be updated with new module path"
        ;;
    *)
        echo -e "${RED}Invalid choice. Exiting.${NC}"
        exit 1
        ;;
esac