#!/bin/sh

set -e

# Installation directory
INSTALL_DIR="${HOME}/.local/bin"
INSTALL_PATH="${INSTALL_DIR}/avds"

# Check if installed
if [ ! -f "$INSTALL_PATH" ]; then
  echo "avds is not installed at $INSTALL_PATH"
  exit 0
fi

# Remove binary
echo "Removing avds..."
rm -f "$INSTALL_PATH"
echo "avds has been removed from $INSTALL_PATH"

# Remind about PATH
echo "Note: Your PATH environment variable was not modified."
echo "If you manually added $INSTALL_DIR to your PATH, you may want to remove it if not needed for other applications."

echo "Uninstallation completed!"