#!/usr/bin/env bash
set -eu -o pipefail

source ./scripts/build/.build_variables.sh
echo "${TARGET}"
go build -o "${TARGET}"  -ldflags "${LDFLAGS}"  "${SOURCE}"
