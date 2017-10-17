#!/usr/bin/env bash
set -eu
source ./scripts/.app_variables.sh

GIT_TAG=$(git tag | awk 'END {print}')
GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
GIT_HASH=$(git rev-parse HEAD | cut -c 1-8)
BUILD_DATE=$(date '+%G-%m-%d')
BUILD_TIME=$(date '+%H:%M:%S')
VERSION_PACKAGE="$APP_SOURCE"



if [[ -z ${app_version+x} ]]
then
    VERSION=$GIT_TAG
else
    VERSION=$app_version
fi


GOOS="${GOOS:-$(go env GOHOSTOS)}"
GOARCH="${GOARCH:-$(go env GOHOSTARCH)}"


export TARGET="build/$APP_NAME-$GOOS-$GOARCH"
export SOURCE="$APP_SOURCE"
export LDFLAGS="\
    -X $VERSION_PACKAGE.APP_NAME=$APP_NAME \
    -X $VERSION_PACKAGE.VERSION=$VERSION \
    -X $VERSION_PACKAGE.BUILD_HASH=$GIT_HASH \
    -X $VERSION_PACKAGE.BUILD_BRANCH=$GIT_BRANCH \
    -X $VERSION_PACKAGE.BUILD_DATE=$BUILD_DATE\
    -X $VERSION_PACKAGE.BUILD_TIME=$BUILD_TIME \
"
