#!/bin/bash

echo Building for all architectures

echo Linux 32bit
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o httpdir32 .
file httpdir32

echo Linux 64bit
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpdir .
file httpdir

echo Linux Arm 32bit
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o httpdir-arm .
file httpdir-arm

echo Linux Arm 64bit
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o httpdir-arm64 .
file httpdir-arm64

echo Windows 32bit
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o httpdir.exe .
file httpdir.exe

echo Windows 64bit
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o httpdir64.exe .
file httpdir64.exe

echo OSX 64bit
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o httpdir-osx .
file httpdir-osx

