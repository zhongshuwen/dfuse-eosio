#!/bin/bash
BASE_DIR="$PWD"
ZSW_LISHI_NODE_BASE_IMAGE_NAME="zhongshuwen/zsw-lishi-node-base"
ZSW_LISHI_NODE_BASE_VERSION="latest"
ZSW_LISHI_NODE_BASE_FULL_NAME="${ZSW_LISHI_NODE_BASE_IMAGE_NAME}:${ZSW_LISHI_NODE_BASE_VERSION}"


buildlishinodebaseimage(){
  docker build -t "${ZSW_LISHI_NODE_BASE_IMAGE_NAME}" . 
  docker tag "${ZSW_LISHI_NODE_BASE_FULL_NAME}" "${ZSW_LISHI_NODE_BASE_IMAGE_NAME}:latest"
}

copyfilefromdockerimage() {
  DOCKER_IMAGE="$1"
  IMAGE_PATH="$2"
  HOST_PATH="$3"
  id=$(docker create $DOCKER_IMAGE)
  docker cp "$id:$IMAGE_PATH" "$HOST_PATH"
  docker rm -v $id
}
pushall(){

  docker push "$ZSW_LISHI_NODE_BASE_FULL_NAME"
  docker push "${ZSW_LISHI_NODE_BASE_IMAGE_NAME}:latest"
}

copybinfiles() {
  copyfilefromdockerimage "$ZSW_LISHI_NODE_BASE_IMAGE_NAME" "/app/zswlishi" "./"
}
buildlishinodebaseimage
copybinfiles

pushall