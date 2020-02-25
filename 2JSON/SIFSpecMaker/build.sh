 #!/bin/bash

VERSION="v0.1.0"

go get

GOARCH=amd64
LDFLAGS="-s -w"
OUT=mkSIFSpecCfg

OUTPATH=./
GOOS="linux" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT