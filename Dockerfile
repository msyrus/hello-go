# Defining App builder image
FROM golang:alpine AS builder

# Add git to determine build git version
RUN apk add --no-cache --update git protoc

# Set apps source directory
ENV SRC_DIR=${GOPATH}/src/github.com/msyrus/hello-go

# Define current working directory
WORKDIR ${SRC_DIR}

COPY go.mod go.sum ./

RUN go mod download

# Copy apps scource code to the image
COPY . .

# Build App
RUN ./build.sh

# Defining App image
FROM alpine:latest

# Copy App binary to image
COPY --from=builder /go/bin/hello /usr/local/bin/hello

EXPOSE 8080

ENTRYPOINT ["hello"]