#!/bin/bash

SCRIPT_DIR="$(dirname "${BASH_SOURCE:-$0}")"
ROOT_DIR=`python -c 'from __future__ import print_function;import os,sys;print(os.path.realpath(sys.argv[1]))' "$SCRIPT_DIR/.."`

cd "$ROOT_DIR/cmd"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -x -o "$ROOT_DIR/out/linux-amd64/v2up" -ldflags "-s -w"
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -x -o "$ROOT_DIR/out/darwin-amd64/v2up" -ldflags "-s -w"

cd "$ROOT_DIR/out/linux-amd64"
tar zcf v2up-linux-amd64.tar.gz v2up
cp v2up-linux-amd64.tar.gz "$ROOT_DIR/out/"

cd "$ROOT_DIR/out/darwin-amd64"
tar zcf v2up-darwin-amd64.tar.gz v2up
cp v2up-darwin-amd64.tar.gz "$ROOT_DIR/out/"

cd -