#!/usr/bin/env bash
set -eu

unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     goos=linux;;
    Darwin*)    goos=darwin;;
    *)          goos=windows
esac

export DOCKER_BUILD_IMAGE="kaleocheng/golang:1.9.0"
export GOOS=$goos
