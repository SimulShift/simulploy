#!/bin/sh

# Echo the current user
echo "Running as user: $(whoami)"

# Start Envoy
envoy -c /etc/envoy/envoy.yaml
