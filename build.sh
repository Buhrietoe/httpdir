#!/bin/bash

echo Building for all architectures

echo Linux 32bit
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags '-w -s' -o target/httpdir32 .
file target/httpdir32

echo Linux 64bit
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o target/httpdir .
file target/httpdir

echo Linux Arm 32bit
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags '-w -s' -o target/httpdir-arm .
file target/httpdir-arm

echo Linux Arm 64bit
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags '-w -s' -o target/httpdir-arm64 .
file target/httpdir-arm64

echo Windows 32bit
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o target/httpdir.exe .
file target/httpdir.exe

echo Windows 64bit
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o target/httpdir64.exe .
file target/httpdir64.exe

echo OSX 64bit
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o target/httpdir-osx .
file target/httpdir-osx


