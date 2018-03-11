#!/usr/bin/env bash

VERSION=$1
USERNAME=jonasfred
IMAGE=systembolaget-article-indexer

echo "Build Docker image"
docker build -t $USERNAME/$IMAGE .
docker tag $USERNAME/$IMAGE:latest $USERNAME/$IMAGE:$version