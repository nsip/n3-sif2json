 #!/bin/bash

### Create toml files at building time ###

SPEC=../SIFSpec/3.4.6.txt
MAKER=../2JSON/SpecCfgMaker
BASE_GO=$MAKER/base-go
BASE_TOML=$MAKER/base-toml
go run var.go main.go -- $SPEC $BASE_GO/config $BASE_TOML/List2JSON $BASE_TOML/Num2JSON $BASE_TOML/Bool2JSON $MAKER/