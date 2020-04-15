#!/bin/sh
echo "Starting"
echo "Installing dependencies..."
go get github.com/tools/godep
echo "Build the project"
go build
