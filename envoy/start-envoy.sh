#!/bin/sh

# Echo the current user
echo "Running as user: $(whoami)"

#!/usr/bin/env sh
set -e

# Define the source and destination directories
SRC_DIR="/etc/letsencrypt/live/backend.simulshift.com"
DEST_DIR="/tmp/letsencrypt/live/backend.simulshift.com"

# Create the destination directory
mkdir -p "$DEST_DIR"

# Copy the certificate files
cp -r "$SRC_DIR"/* "$DEST_DIR"/

# Change the ownership of the copied files
chown -R 101:101 "$DEST_DIR"

# Execute Envoy
exec envoy -c /etc/envoy/envoy.yaml