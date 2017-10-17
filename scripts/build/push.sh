#!/usr/bin/env bash
set -eu -o pipefail

source ./scripts/build/.git_variables.sh

docker login -u $DOCKER_USER -p $DOCKER_PASS
docker push kaleocheng/wormhole:$GIT_VERSION
docker push kaleocheng/wormhole:latest

