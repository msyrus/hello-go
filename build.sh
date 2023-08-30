#! /bin/sh
set -e

if ! [ -x "$(command -v go)" ]; then
    echo "go is not installed"
    exit
fi
if ! [ -x "$(command -v git)" ]; then
    echo "git is not installed"
    exit
fi
if ! [ -x "$(command -v protoc)" ]; then
    echo "protoc is not installed"
    exit
fi

GOBIN="$(go env GOPATH)/bin"
PATH="${PATH}:${GOBIN}"

if ! [ -x "$GOBIN/protoc-gen-go" ]; then
    echo "Installing protoc-gen-go ..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31
fi

if ! [ -x "$GOBIN/protoc-gen-go-grpc" ]; then
    echo "Installing protoc-gen-go-grpc ..."
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3
fi

# Prepare binary version
COMMIT=`git rev-parse --short HEAD`
TAG=$(git describe --exact-match --abbrev=0 --tags ${COMMIT} 2> /dev/null || true)

if [ -z "${TAG}" ]; then
    VERSION=${COMMIT}
else
    VERSION=${TAG}
fi

if [ -n "$(git diff --shortstat 2> /dev/null | tail -n1)" ]; then
    VERSION="${VERSION}-dirty"
fi

# Compiling protobuf
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false,paths=source_relative \
    proto/**/*.proto

# Building go binary
go install -v -ldflags="-X github.com/msyrus/hello-go/version.Version=${VERSION}" ./cmd/...