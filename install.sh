#!/bin/bash

# SyntaxRush Quick Install Script
# This script provides easy installation options for SyntaxRush

set -e

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${BLUE}🚀 SyntaxRush Installation${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "Choose installation method:"
echo ""
echo "1. Global Installation (Recommended)"
echo "   - Install syntaxrush globally to /usr/local/bin"
echo "   - Available from anywhere: 'syntaxrush practice file.go'"
echo ""
echo "2. Shell Completion Setup"
echo "   - Add auto-completion for bash/zsh/fish"
echo "   - Enables tab completion for commands and files"
echo ""
echo "3. Development Build"
echo "   - Build locally for development"
echo "   - Run with './syntaxrush'"
echo ""

read -p "Enter your choice (1-3): " choice

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
        ;;
    *)
        echo -e "${YELLOW}Invalid choice. Exiting.${NC}"
        exit 1
        ;;
esac