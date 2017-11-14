#!/bin/bash

echo Building...

CGO_ENABLED=0 go build -v -ldflags '-w -s' -o httpdir .
file httpdir
