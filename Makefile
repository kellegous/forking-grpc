export GOPATH := $(shell pwd)
export PATH := $(shell pwd)/bin:$(PATH)

ALL: bin/be bin/fe php/Pkg/MaffClient.php

bin/protoc-gen-go:
	go install ./src/maff/vendor/github.com/golang/protobuf/protoc-gen-go

src/maff/pkg/maff.pb.go: bin/protoc-gen-go maff.proto
	protoc \
		-I . \
		maff.proto \
		--go_out=plugins=grpc:./src/maff/pkg

bin/fe: src/maff/pkg/maff.pb.go $(shell find src -type f)
	go install maff/cmd/fe

bin/be: src/maff/pkg/maff.pb.go $(shell find src -type f)
	go install maff/cmd/be

php/Pkg/MaffClient.php: maff.proto
	mkdir -p php
	protoc --proto_path=. --php_out=php --grpc_out=php \
		--plugin=protoc-gen-grpc=grpc/bins/opt/grpc_php_plugin \
		maff.proto

build-grpc:
	script/build-grpc

test: bin/be build-grpc php/Pkg/MaffClient.php
	script/run-test