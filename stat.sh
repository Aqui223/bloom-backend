#!/bin/sh

# Directories to scan
DIRS="internal cmd"

# Count .go files
FILES_COUNT=$(find $DIRS -type f -name "*.go" 2>/dev/null | wc -l)

# Count folders (unique directories that contain .go files)
FOLDERS_COUNT=$(find $DIRS -type f -name "*.go" 2>/dev/null -exec dirname {} \; | sort -u | wc -l)

# Count total lines in .go files
LINES_COUNT=$(find $DIRS -type f -name "*.go" 2>/dev/null -exec cat {} \; | wc -l)

echo "Lines: $LINES_COUNT"
echo "Files: $FILES_COUNT"
echo "Folders: $FOLDERS_COUNT"
