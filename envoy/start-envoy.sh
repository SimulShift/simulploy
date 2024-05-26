#!/bin/sh

# Echo the current user
echo "Running as user: $(whoami)"

#!/usr/bin/env sh
set -e

# Define the source and destination directories
SRC_DIR="/etc/letsencrypt"
DEST_DIR="/tmp/letsencrypt"

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

# List the copied directory and its contents after changing ownership
echo "Files in destination directory ($DEST_DIR) after changing ownership:"
ls -lR "$DEST_DIR"

# print arguments
echo "Arguments: $*"

# Execute the command as envoy user
su-exec envoy "$@"
