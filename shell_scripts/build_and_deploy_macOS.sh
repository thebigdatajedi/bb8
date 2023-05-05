#!/bin/bash

# This script is used to build and install the latest version of bb8.
# Intended to be run on macOS.
clear

# params
echo "$(date) - setting params"
source_binary="bb8"
target_directory="/usr/local/bin"

# check if the source binary exists on the root dir and if it does then delete - clean step
echo "$(date) - clean step started"
if [ -f $source_binary ]; then
    rm $source_binary
    echo "$(date) - stale binary deleted - clean step complete"
fi
echo "$(date) - clean step complete"

# Build
echo "$(date) - build started"
go build bb8.go
echo "$(date) - build complete"

# Install
# This code is used to get the current working directory of the script,
# with that I can then bootstrap a relative path to the source binary, other wise bash didn't know what the
# .. in the ../$source_binary parameter meant.
# I corrected this by creating dynamic context of the current working directory of the script.
echo "$(date) - bootstrapping working dir started"
working_dir=$( cd "$(dirname "${BASH_SOURCE[0]}")" || exit ; pwd -P )
cd "$working_dir" || exit
echo "$(date) - bootstrapping working dir complete"

echo "$(date) - deploy started"
./macOS_deploy.sh "$working_dir/../$source_binary" $target_directory $source_binary
echo "$(date) - deploy complete"
# end of script