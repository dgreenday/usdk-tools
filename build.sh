#!/bin/bash

set -x
set -e

COPY_DIR=`pwd`/build/src/launchpad.net/ubuntu-sdk-tools
BUILD_DIR=`pwd`/build

mkdir -p $COPY_DIR

rm -f $BUILD_DIR/usdk-target
rm -f $BUILD_DIR/usdk-wrapper
rm -f $COPY_DIR/* -r

rsync -av --progress . $COPY_DIR --exclude build

export GOPATH=`pwd`/build/
go get github.com/tools/godep

cd $COPY_DIR
$GOPATH/bin/godep go build -o $BUILD_DIR/usdk-target ./usdk-target/
$GOPATH/bin/godep go build -o $BUILD_DIR/usdk-wrapper ./usdk-wrapper/
