#!/bin/bash

# Build
docker build -t inventory .

# Tag
docker tag inventory asia-southeast1-docker.pkg.dev/inventory-476214/inventory/api:latest 

# Push
docker push asia-southeast1-docker.pkg.dev/inventory-476214/inventory/api:latest