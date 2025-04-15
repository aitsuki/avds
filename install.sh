#!/bin/sh

set -e

# Determine OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case $OS in
  darwin) OS="darwin" ;;
  linux) OS="linux" ;;
  *) echo "Error: Unsupported operating system: $OS"; exit 1 ;;
esac

# Determine architecture (only support 64-bit)
ARCH=$(uname -m)
case $ARCH in
  x86_64) ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *) echo "Error: Only support amd64/arm64"; exit 1 ;;
esac

# Get latest version
echo "Fetching latest release..."
LATEST_VERSION=$(curl -s "https://api.github.com/repos/aitsuki/avds/releases/latest" | grep -o '"tag_name": "v[^"]*' | cut -d'"' -f4)
if [ -z "$LATEST_VERSION" ]; then
  echo "Error: Unable to get latest version"
  exit 1
fi

# Download URL
DOWNLOAD_URL="https://github.com/aitsuki/avds/releases/download/${LATEST_VERSION}/avds-${LATEST_VERSION}-${OS}-${ARCH}"

# Installation directory
INSTALL_DIR="${HOME}/.local/bin"
mkdir -p "$INSTALL_DIR"
INSTALL_PATH="${INSTALL_DIR}/avds"

# Download and install
echo "Downloading avds ${LATEST_VERSION} (${OS}/${ARCH})..."
curl -sL "$DOWNLOAD_URL" -o "$INSTALL_PATH"
chmod +x "$INSTALL_PATH"

echo "avds has been installed to $INSTALL_PATH"

# Check PATH
if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
  echo "Please add $INSTALL_DIR to your PATH environment variable"
  case "$SHELL" in
    */bash*)
      echo "You can add it to ~/.bashrc by running:"
      echo "  echo 'export PATH=\"\$PATH:$INSTALL_DIR\"' >> ~/.bashrc"
      echo "  source ~/.bashrc"
      ;;
    */zsh*)
      echo "You can add it to ~/.zshrc by running:"
      echo "  echo 'export PATH=\"\$PATH:$INSTALL_DIR\"' >> ~/.zshrc"
      echo "  source ~/.zshrc"
      ;;
  esac
fi

echo "Installation completed! Run 'avds' to start"