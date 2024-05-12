#!/bin/bash

# Initialize the variable to an empty string
d=""

# Check if the "-d" flag is provided
if [[ "$*" == *"-d"* ]]; then
    # If the "-d" flag is provided, set the variable to "-d"
    d="-d"
fi

# Execute the docker-compose commands with the "-d" flag if needed
docker compose down postgres-prod
docker volume rm docker_postgres-data
docker compose up $d postgres-prod

