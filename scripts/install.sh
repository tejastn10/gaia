#!/usr/bin/env bash

set -e

# Determine OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Set binary name based on OS and ARCH
if [[ "$OS" == "linux" ]]; then
  BINARY="gaia-linux-amd64"
elif [[ "$OS" == "darwin" ]]; then
  BINARY="gaia-darwin-arm64"
elif [[ "$OS" == "msys" || "$OS" == "cygwin" ]]; then
  BINARY="gaia-windows-amd64.exe"
else
  echo "Unsupported OS: $OS"
  exit 1
fi

VERSION=$(curl -s https://api.github.com/repos/tejastn10/gaia/releases/latest | jq -r '.tag_name')
BASE_URL="https://github.com/tejastn10/gaia/releases/download/$VERSION"
FILE="$BINARY.zip"

# Download the binary
echo "Downloading $BINARY from $BASE_URL/$FILE ..."
curl -LO "$BASE_URL/$FILE"

# Extract if on Unix-like system
if [[ "$OS" != "msys" && "$OS" != "cygwin" ]]; then
  unzip "$FILE" -d gaia_temp
  BINARY_PATH=$(find gaia_temp -type f -name "$BINARY")
else
  # Windows users need a different extraction method
  BINARY_PATH="$BINARY"
fi

# Install: move binary to /usr/local/bin (requires sudo)
echo "Installing gaia to /usr/local/bin ..."
sudo mv "$BINARY_PATH" /usr/local/bin/gaia
sudo chmod +x /usr/local/bin/gaia

# Cleanup
rm -rf "$FILE" gaia_temp

echo "Gaia installed successfully!"
