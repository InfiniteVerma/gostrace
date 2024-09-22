#!/bin/bash

for filename in $PWD/*.cpp; do
    echo Compiling: $filename
    cpp_file_name=$(echo $filename | sed -E 's/.*example\///g')
    obj_name=$(echo $cpp_file_name | sed -E 's/\.cpp//')

    g++ $cpp_file_name -o $obj_name
done
