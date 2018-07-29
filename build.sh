#! /bin/sh
set -e

PROJ="hello-go"
ORG_PATH="github.com/msyrus"
REPO_PATH="${ORG_PATH}/${PROJ}"

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

if [ -z "${GOPATH}" ]; then
    echo "set GOPATH"
    exit
fi

PATH="${PATH}:${GOPATH}/bin"

if ! [ -x "$GOPATH/bin/dep" ]; then
    echo "Installing dep ..."
    go get -u github.com/golang/dep/cmd/dep
fi

if ! [ -x "$GOPATH/bin/protoc-gen-go" ]; then
    echo "Installing protoc-gen-go ..."
    go get -u github.com/golang/protobuf/protoc-gen-go
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
for d in proto/* ; do
    echo "Compiling $d";
    protoc --go_out=plugins=grpc:"$GOPATH/src" $d/*.proto
done

# Building go binary
dep ensure -v -vendor-only
go install -v -ldflags="-X ${REPO_PATH}/version.Version=${VERSION}" ./cmd/...