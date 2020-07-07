#!/usr/bin/env bash
. ./build/docker.env

VERSION=$(git rev-parse HEAD)
# VERSION="PROD_2020061901"
IMAGE_NAME="gcr.io/${PROJECT_NAME}/${SERVICE_NAME}"
docker build . -f ./build/Dockerfile -t "${IMAGE_NAME}:${VERSION}" -t "${IMAGE_NAME}:latest"
gcloud auth configure-docker --quiet
docker push "${IMAGE_NAME}:latest"
docker push "${IMAGE_NAME}:${VERSION}"
