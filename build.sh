#!/bin/bash

set -e

VERSION=$(jq -r '.version' manifest.json)
GOOS=linux GOARCH=arm GOARM=7 go build --ldflags="-X main.version=$VERSION" -o jetkvm-plugin-serialkvm .
tar -czvf serialkvm.tar.gz manifest.json jetkvm-plugin-serialkvm