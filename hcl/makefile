ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
GOOS=darwin
GOARCH=amd64


build:
	cd ./lib/math && cargo build --release
	cp ./lib/math/target/release/libmath.a ./lib
	go build -ldflags="-r $(ROOT_DIR)lib" main.go