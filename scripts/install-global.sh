#!/bin/bash

# SyntaxRush Global Installation Script
# This script installs SyntaxRush globally so you can use 'syntaxrush' from anywhere

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
BINARY_NAME="syntaxrush"
INSTALL_DIR="/usr/local/bin"
ASSETS_DIR="/usr/local/share/syntaxrush/assets"
BUILD_DIR="$(pwd)"

echo -e "${BLUE}🚀 SyntaxRush Global Installation${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is not installed. Please install Go first.${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Go found: $(go version)${NC}"

# Check if we're in the right directory
if [ ! -f "go.mod" ] || [ ! -f "main.go" ]; then
    echo -e "${RED}❌ Please run this script from the SyntaxRush root directory${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Found SyntaxRush project files${NC}"

# Build the binary
echo -e "${BLUE}🔨 Building SyntaxRush...${NC}"
go build -o "${BINARY_NAME}" -ldflags="-s -w" .

if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Build failed${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Build successful${NC}"

# Check if install directory exists and is writable
if [ ! -d "$INSTALL_DIR" ]; then
    echo -e "${YELLOW}⚠️  Creating install directory: $INSTALL_DIR${NC}"
    sudo mkdir -p "$INSTALL_DIR"
fi

# Install the binary
echo -e "${BLUE}📦 Installing binary to $INSTALL_DIR...${NC}"

if [ -w "$INSTALL_DIR" ]; then
    # Can write without sudo
    cp "$BINARY_NAME" "$INSTALL_DIR/"
    chmod +x "$INSTALL_DIR/$BINARY_NAME"
else
    # Need sudo
    echo -e "${YELLOW}🔒 Requesting sudo access to install to $INSTALL_DIR${NC}"
    sudo cp "$BINARY_NAME" "$INSTALL_DIR/"
    sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"
fi

# Install assets
if [ -d "assets" ]; then
    echo -e "${BLUE}📁 Installing assets to $ASSETS_DIR...${NC}"
    
    if [ -w "$(dirname "$ASSETS_DIR")" ] 2>/dev/null; then
        # Can write without sudo
        mkdir -p "$ASSETS_DIR"
        cp -r assets/* "$ASSETS_DIR/"
    else
        # Need sudo
        echo -e "${YELLOW}🔒 Requesting sudo access to install assets${NC}"
        sudo mkdir -p "$ASSETS_DIR"
        sudo cp -r assets/* "$ASSETS_DIR/"
    fi
    
    echo -e "${GREEN}✅ Assets installed${NC}"
else
    echo -e "${YELLOW}⚠️  No assets directory found, skipping asset installation${NC}"
fi

# Clean up build artifact
rm "$BINARY_NAME"

# Verify installation
if command -v "$BINARY_NAME" &> /dev/null; then
    echo -e "${GREEN}✅ Installation successful!${NC}"
    echo ""
    echo -e "${BLUE}🎉 SyntaxRush is now globally available!${NC}"
    echo ""
    echo "Usage examples:"
    echo "  ${BINARY_NAME} --help                    # Show help"
    echo "  ${BINARY_NAME} practice                  # Practice with sample files"
    echo "  ${BINARY_NAME} practice main.go          # Practice with your file"
    echo "  ${BINARY_NAME} practice go --quick       # Quick practice with Go sample"
    echo "  ${BINARY_NAME} practice python --mute    # Practice Python silently"
    echo "  ${BINARY_NAME} stats                     # View your statistics"
    echo "  ${BINARY_NAME} version                   # Show version info"
    echo ""
    echo -e "${GREEN}Current version: $($BINARY_NAME version --short 2>/dev/null || echo 'v1.0.0')${NC}"
else
    echo -e "${RED}❌ Installation failed - binary not found in PATH${NC}"
    echo -e "${YELLOW}💡 You may need to add $INSTALL_DIR to your PATH${NC}"
    echo ""
    echo "Add this to your ~/.bashrc or ~/.zshrc:"
    echo "export PATH=\"\$PATH:$INSTALL_DIR\""
    exit 1
fi

echo ""
echo -e "${BLUE}🔧 To uninstall later, run:${NC}"
echo "sudo rm $INSTALL_DIR/$BINARY_NAME"
echo "sudo rm -rf $ASSETS_DIR"
