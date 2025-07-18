#!/bin/bash

echo "🔧 SyntaxRush Sound Test"
echo "========================"
echo ""
echo "Testing different sound methods on your system:"
echo ""

echo "1. Testing terminal bell..."
echo -e "\a"
sleep 1

echo "2. Testing system beep commands..."
if command -v paplay &> /dev/null; then
    echo "   Trying paplay..."
    paplay /usr/share/sounds/alsa/Front_Left.wav 2>/dev/null || echo "   paplay failed"
fi

if command -v aplay &> /dev/null; then
    echo "   Trying aplay..."
    aplay /usr/share/sounds/alsa/Front_Left.wav 2>/dev/null || echo "   aplay failed"
fi

if command -v speaker-test &> /dev/null; then
    echo "   Trying speaker-test..."
    timeout 0.5 speaker-test -t sine -f 800 -l 1 2>/dev/null || echo "   speaker-test failed"
fi

echo ""
echo "3. If you heard any sound, the error beep should work in SyntaxRush."
echo "4. If no sound was heard, your system might have:"
echo "   - Audio disabled"
echo "   - Terminal bell disabled"
echo "   - Missing audio drivers"
echo ""
echo "🚀 Now test SyntaxRush:"
echo "   1. Run: ./syntaxrush"
echo "   2. Press Enter to start typing"
echo "   3. Type a wrong character (e.g., 'x' instead of 'p' for 'package')"
echo "   4. You should see '❌ MISTAKE! ❌' message and hear a beep"
echo ""
