#!/usr/bin/env bash
set -eu -o pipefail

source ./scripts/build/.git_variables.sh

docker build -t kaleocheng/wormhole:$GIT_VERSION .
docker tag kaleocheng/wormhole:$GIT_VERSION kaleocheng/wormhole:latest

