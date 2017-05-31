#!/bin/bash

# Create a go workspace in a container, and run any command
#
# $PWD is mounted into the correct go workspace loaction
# All output artifacts are placed in $PWD/_docker_workspace
#

set -e
set -x

CURDIR="$(dirname $BASH_SOURCE)"

. $CURDIR/_build-lib.sh

: ${build_img:=${BUILD_IMG_TAG}}

# Need to make the directory before docker, to keep it owned by local user
srcdir=src/github.com/F5Networks/k8s-bigip-ctlr
wkspace=${PWD}/_docker_workspace
mkdir -p $wkspace/$srcdir

RUN_ARGS=( \
  --rm
  -v $wkspace:/build
  -v $PWD:/build/$srcdir:ro
  --workdir  /build/$srcdir
  -e GOPATH=/build
  -e CLEAN_BUILD="${CLEAN_BUILD}"
  -e IMG_TAG="${IMG_TAG}"
  -e BUILD_IMG_TAG="${BUILD_IMG_TAG}"
  -e LOCAL_USER_ID=$(id -u)
)

# Add -it if caller is a terminal
if [ -t 0 ]; then
  RUN_ARGS+=( "-it" )
fi

# Run the user provided args
docker run "${RUN_ARGS[@]}" "$build_img" "$@"
