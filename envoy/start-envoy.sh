#!/bin/sh

# Echo the current user
echo "Running as user: $(whoami)"

#!/usr/bin/env sh
set -e

# Define the source and destination directories
SRC_DIR="/etc/letsencrypt/live/backend.simulshift.com"
DEST_DIR="/tmp/letsencrypt/live/backend.simulshift.com"

# Print files in source directory before copying
echo "Files in source directory ($SRC_DIR) before copying:"
ls -l "$SRC_DIR"

# Create the destination directory
mkdir -p "$DEST_DIR"

# Copy the certificate files
cp -r "$SRC_DIR"/* "$DEST_DIR"/

# Change the ownership of the copied files
chown -R 101:101 "$DEST_DIR"

# List the copied files after changing ownership
echo "Files in destination directory ($DEST_DIR) after copying:"
ls -l "$DEST_DIR"

# Check if privkey.pem exists in the destination directory
PRIVATE_KEY="$DEST_DIR/privkey.pem"
if [ -f "$PRIVATE_KEY" ]; then
    echo "Privkey file exists: $PRIVATE_KEY"
    # Check permissions and ownership of the private key file
    echo "Checking permissions and ownership of $PRIVATE_KEY:"
    ls -l "$PRIVATE_KEY"
    echo "Running 'stat' command on $PRIVATE_KEY:"
    stat "$PRIVATE_KEY"
    # Print the content of the private key file
    echo "Content of $PRIVATE_KEY:"
    cat "$PRIVATE_KEY"
else
    echo "Error: Privkey file does not exist: $PRIVATE_KEY"
fi

# print arguments
echo "Arguments: $*"

su-exec envoy "$@"
