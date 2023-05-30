#!/bin/bash

VER=0.1.8
GIT_ID=$(git rev-parse HEAD | cut -c1-7)
GO_VER=$(go version | cut -d' ' -f3)
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_TIME=$(date '+%Y-%m-%d-%H:%M:%S')
BUILD_FLAGS=''

if [ -z "$LINUX" ]; then
    if [ -n "$STATIC_BUILD" ]; then
        BUILD_FLAGS="$BUILD_FLAGS -a"
        CGO_ENABLED=0 go build $BUILD_FLAGS -ldflags "-X github.com/funkygao/golib/version.BuildDate=$BUILD_TIME -X github.com/funkygao/golib/version.GoVersion=$GO_VER -X github.com/funkygao/golib/version.Version=$VER -X github.com/funkygao/golib/version.Revision=${GIT_ID}${GIT_DIRTY}" ./cmd/flexdb
    else
        go build $BUILD_FLAGS -ldflags "-X github.com/funkygao/golib/version.BuildDate=$BUILD_TIME -X github.com/funkygao/golib/version.GoVersion=$GO_VER -X github.com/funkygao/golib/version.Version=$VER -X github.com/funkygao/golib/version.Revision=${GIT_ID}${GIT_DIRTY}" ./cmd/flexdb
    fi
else
    echo "building for Linux"
    GOOS=linux GOARCH=amd64 go build $BUILD_FLAGS -ldflags "-X github.com/funkygao/golib/version.BuildDate=$BUILD_TIME -X github.com/funkygao/golib/version.GoVersion=$GO_VER -X github.com/funkygao/golib/version.Version=$VER -X github.com/funkygao/golib/version.Revision=${GIT_ID}${GIT_DIRTY}" ./cmd/flexdb
fi

