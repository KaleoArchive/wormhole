#!/usr/bin/env bash
set -eu
source ./scripts/.app_variables.sh
source ./scripts/.docker_variables.sh


command="bash"
if ! [[ -z "$@" ]]
then
  command=$@
fi

echo "Make with Docker..."
docker run --rm -ti \
-e GOOS=$GOOS \
-w /go/src/$APP_SOURCE \
-v `pwd`:/go/src/$APP_SOURCE \
$DOCKER_BUILD_IMAGE \
$command
