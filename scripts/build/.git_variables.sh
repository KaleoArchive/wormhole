
#!/bin/bash -e

# parse the current git commit hash
COMMIT=`git rev-parse HEAD | cut -c 1-8`

# check if the current commit has a matching tag
TAG=$(git describe --exact-match --abbrev=0 --tags ${COMMIT} 2> /dev/null || true)

# use the matching tag as the version, if available
if [ -z "$TAG" ]; then
    VERSION=$COMMIT
else
    VERSION=$TAG
fi

# check for changed files (not untracked files)
if [ -n "$(git diff --shortstat 2> /dev/null | tail -n1)" ]; then
    VERSION="${VERSION}-dirty"
fi


BRANCH=$(git rev-parse --abbrev-ref HEAD)



export GIT_VERSION=$VERSION
export GIT_COMMIT=$COMMIT
export GIT_TAG=$TAG
export GIT_BRANCH=$BRANCH
