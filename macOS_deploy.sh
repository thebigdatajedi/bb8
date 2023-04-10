#!/bin/bash
echo "script parameters: $*"

# get the first parameter
source_binary=$1
target_dir=$2

# check if the source binary already exists on target directory
# then delete it
if [ -f $target_dir/$source_binary ]; then
    rm $target_dir/$source_binary
    echo "source binary deleted from target directory"
fi

# move the source binary to target directory
mv $source_binary $target_dir

# check if the source binary exists on target directory
# then print the result and throw an error
if [ -f $target_dir/$source_binary ]; then
    echo "source binary moved to target directory"
else
    echo "source binary not moved to target directory"
    exit 1
fi

# end of script
