#!/bin/bash

# get the first parameter
echo "$(date) - setting $1 $2 $3"
source_binary=$1
target_dir=$2
simple_source_binary_name=$3

# check if the source binary already exists on target directory
# then delete it
echo "$(date) - checking if source binary exists on target directory"
if [ -f "$target_dir"/"$simple_source_binary_name" ]; then
    rm "$target_dir"/"$simple_source_binary_name"
    echo "$(date) - source binary deleted from target directory"
fi

# first if in nested if statements
#if target directory exists
echo "$(date) - checking if target directory exists && source binary exists"
if [ -d "$target_dir" ]; then
  #and if source binary exists
    if [ -f "$source_binary" ]; then
        #then move the source binary to target directory
        mv "$source_binary" "$target_dir"
    #else message user
    else
        echo "$(date) - $source_binary does not exist"
        exit 1
    fi
#else message user
else
    echo "$(date) - $target_dir does not exist"
    exit 1
fi

# check if the source binary exists on target directory
# then print the result and throw an error
if [ -f "$target_dir"/"$simple_source_binary_name" ]; then
    echo "$(date) - source binary moved to target directory successfully"
else
    echo "$(date) - source binary failed to move to target directory"
    exit 1
fi
# end of script