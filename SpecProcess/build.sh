 #!/bin/bash

### Create toml files at building time ###

if [ $# -ne 1 ]; then
    echo "Input SIF(txt) Spec File Path"
    exit -1
fi

if [ ! -f $1 ]; then
    echo "SIF txt Spec does not exist"
    exit -1
fi

SPEC=$1
MAKER=../2JSON/SpecCfgMaker
BASE_GO=$MAKER/base-go
BASE_TOML=$MAKER/base-toml
CGO_ENABLED=0 go run var.go main.go -- $SPEC $BASE_GO/spec $BASE_TOML/List2JSON $BASE_TOML/Num2JSON $BASE_TOML/Bool2JSON $MAKER/