CGO_CFLAGS_ALLOW ?= "^.*$$"

all: build

build:
	CGO_CFLAGS_ALLOW=$(CGO_CFLAGS_ALLOW) go build

install:
	CGO_CFLAGS_ALLOW=$(CGO_CFLAGS_ALLOW) go install
