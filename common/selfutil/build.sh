#!/bin/bash

go build -buildmode=c-shared -o output/macos/libadd.dylib add_lib.go

