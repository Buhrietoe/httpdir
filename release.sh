#!/bin/bash

echo Building for all architectures

echo Linux 32bit
TARGET=target/linux-32/httpdir
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags '-w -s' -o $TARGET .
file $TARGET
echo
echo Linux 64bit
TARGET=target/linux-64/httpdir
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o $TARGET .
file $TARGET
echo
echo Linux Arm 32bit
TARGET=target/linux-arm-32/httpdir
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags '-w -s' -o $TARGET .
file $TARGET
echo
echo Linux Arm 64bit
TARGET=target/linux-arm-64/httpdir
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags '-w -s' -o $TARGET .
file $TARGET
echo
echo Linux MIPSLE 32bit
TARGET=target/linux-mipsle-32/httpdir
CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -ldflags '-w -s' -o $TARGET .
file $TARGET
echo
echo Windows 32bit
TARGET=target/linux-windows-32/httpdir.exe
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o $TARGET .
file $TARGET
echo
echo Windows 64bit
TARGET=target/linux-windows-64/httpdir.exe
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $TARGET .
file $TARGET
echo
echo OSX 64bit
TARGET=target/linux-osx-64/httpdir
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $TARGET .
file $TARGET

