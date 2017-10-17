#!/usr/bin/env bash
set -eu -o pipefail

BUILDDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

jobs=(
    "GOOS=darwin GOARCH=amd64   $BUILDDIR/binary.sh" \
    "GOOS=linux GOARCH=amd64   $BUILDDIR/binary.sh" \
    "GOOS=windows GOARCH=amd64 $BUILDDIR/binary.sh" \
)

echo "Building binaries for Linux MacOS and Windows platforms"
parallel --no-notice ::: "${jobs[@]}"
exit 0
