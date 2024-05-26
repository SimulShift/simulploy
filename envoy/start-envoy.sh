#!/bin/sh

# Echo the current user
echo "Running as user: $(whoami)"

#!/usr/bin/env sh
set -e

# Define the source and destination directories
SRC_DIR="/etc/letsencrypt"
DEST_DIR="/tmp"

# Print files in source directory before copying
echo "Files in source directory ($SRC_DIR) before copying:"
ls -l "$SRC_DIR"

# Create the destination directory
mkdir -p "$DEST_DIR"

# Print files in destination directory after creation
echo "Files in destination directory ($DEST_DIR) after creation:"
ls -l "$DEST_DIR"

# Copy the entire letsencrypt directory recursively
cp -r "$SRC_DIR" "$DEST_DIR"

# Print files in destination directory after copying
echo "Files in destination directory ($DEST_DIR) after copying:"
ls -l "$DEST_DIR"

# Change the ownership of the copied directory and its contents
chown -R 101:101 "$DEST_DIR"

# print arguments
echo "Arguments: $*"

# Define the envoy user
ENVOY_USER="envoy"

# Check what user envoy is set to
ENVOY_INFO=$(getent passwd "$ENVOY_USER" || true)

# Print information about the envoy user
echo "Information about the envoy user:"
echo "$ENVOY_INFO"

# Extract UID and GID from the envoy user information
ENVOY_UID=$(echo "$ENVOY_INFO" | cut -d: -f3)
ENVOY_GID=$(echo "$ENVOY_INFO" | cut -d: -f4)

echo "envoy user UID: $ENVOY_UID"
echo "envoy user GID: $ENVOY_GID"

# Execute the command as envoy user
su-exec envoy "$@"
