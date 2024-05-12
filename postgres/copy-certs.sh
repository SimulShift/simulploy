#!/bin/bash

CERT_PATH="/etc/letsencrypt/live/backend.simulshift.com"

# Copy the certificates, following symbolic links (-L option)
cp -L $CERT_PATH/fullchain.pem $CERT_PATH/fullchain-docker.pem
cp -L $CERT_PATH/privkey.pem $CERT_PATH/privkey-docker.pem

# Set permissions to be read-only for the user
chmod 400 $CERT_PATH/fullchain-docker.pem
chmod 400 $CERT_PATH/privkey-docker.pem
chown 999 $CERT_PATH/privkey-docker.pem

echo "Certificates copied and permissions set."

