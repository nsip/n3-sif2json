#!/bin/bash

# delete all binary files
find . -type f -executable -exec sh -c "file -i '{}' | grep -q 'x-executable; charset=binary'" \; -print | xargs rm -f

rm -rf ./build
rm -f *.log ./config/*.log
rm -f ./data/*.json
rm -f *.xml *.json

