#!/usr/bin/env bash

rm -rf ./example.app
rice clean
mkdir -p example.app/Contents/MacOS
cp ./docs/info.plist ./example.app/Contents/Info.plist
rice embed-go
go build -o example.app/Contents/MacOS/example
open example.app