#!/bin/bash

# SyntaxRush Build Script
echo "🧠⚡ Building SyntaxRush..."

# Clean previous build
if [ -f "syntaxrush" ]; then
    rm syntaxrush
    echo "🧹 Cleaned previous build"
fi

# Build the application
echo "🔨 Compiling Go code..."
go build -o syntaxrush

if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
    echo ""
    echo "📦 Binary created: ./syntaxrush"
    echo "📊 Size: $(du -h syntaxrush | cut -f1)"
    echo ""
    echo "🎯 Usage:"
    echo "  ./syntaxrush                    # Start with sample code"
    echo "  ./syntaxrush path/to/file.go    # Start with specific file"
    echo ""
    echo "Practice typing real code. Master syntax! 🧠⚡"
else
    echo "❌ Build failed!"
    exit 1
fi
