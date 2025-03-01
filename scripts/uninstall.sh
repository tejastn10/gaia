#!/usr/bin/env bash

set -e

echo "Removing Gaia..."
sudo rm -f /usr/local/bin/gaia
rm -rf ~/.gaia

echo "Gaia has been successfully uninstalled."
