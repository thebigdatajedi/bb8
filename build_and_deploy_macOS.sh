#!/bin/bash

# This script is used to build and install the latest version of bb8.
# Intended to be run on macOS.

# Build
go build bb8.go

# Install
./macOS_deploy.sh bb8 /usr/local/bin/
