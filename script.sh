#!/bin/bash

# Remove the old container if it exists
docker rm -f forum-container 2>/dev/null

# # Remove the old image if it exists
docker rmi -f forum-app 2>/dev/null

# Build the new image
docker build -t forum-app -f build/Dockerfile .

# Run the new container
docker run -d -p 8080:8080 --name forum-container forum-app

# Show logs
docker logs -f forum-container