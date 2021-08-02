#/bin/bash

# ini silahkan diganti dengan username di dockerhub
USERNAME=stalchmst

docker build -t $USERNAME/alta-docker:latest .
docker push $USERNAME/alta-docker:latest

docker run -d --name altaContainer -p 3000:8080 $USERNAME/alta-docker:latest
