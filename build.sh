#!/bin/bash
# Build script per moon-phase
# Usa Clang invece di GCC (che crasha con GTK CGO)

set -e

echo "Building moon-phase con Clang..."
CC=clang CGO_ENABLED=1 go build -o moon-phase

if [ -f moon-phase ]; then
    echo "✓ Compilazione completata!"
    ls -lh moon-phase
    echo ""
    echo "Esegui con: ./moon-phase"
else
    echo "✗ Errore di compilazione"
    exit 1
fi
