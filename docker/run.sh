#!/bin/bash

echo "Run and create container (image should be exist)"
docker run -d --name altaContainer -p 3000:8080 alta-docker:latest

# docker run -d --name <container-name> \
#   -p <host-port>:<container-port> \
#   -p <host-port>:<container-port> \
#   -e <ENVNAME=VALUE> \
#   -e <ENVNAME=VALUE> \
#   <image-name>