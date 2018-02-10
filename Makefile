CGO_CFLAGS_ALLOW ?= "^.*$$"

all: proto build

proto:
	protoc -I rc/ rc/rc.proto --go_out=plugins=grpc:rc

build:
	CGO_CFLAGS_ALLOW=$(CGO_CFLAGS_ALLOW) go build

install:
	CGO_CFLAGS_ALLOW=$(CGO_CFLAGS_ALLOW) go install
