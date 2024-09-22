#!/bin/bash

for filename in $PWD/*.cpp; do
    echo Cleaning obj of file: $filename
    obj_name=$(echo $filename | sed -E 's/.*example\///g' | sed -E 's/\.cpp//')
    rm $obj_name
done
