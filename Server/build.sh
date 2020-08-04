 #!/bin/bash

set -e

rm -rf ./build

go get

GOARCH=amd64
LDFLAGS="-s -w"
OUT=server

OUTPATH=./build/linux64/
mkdir -p $OUTPATH
CGO_ENABLED=0 GOOS="linux" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT
mv $OUT $OUTPATH
cp ./*.toml $OUTPATH

# OUTPATH=./build/win64/
# mkdir -p $OUTPATH
# CGO_ENABLED=0 GOOS="windows" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT.exe
# mv $OUT.exe $OUTPATH
# cp ./*.toml $OUTPATH

# OUTPATH=./build/mac/
# mkdir -p $OUTPATH
# CGO_ENABLED=0 GOOS="darwin" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT
# mv $OUT $OUTPATH
# cp ./*.toml $OUTPATH

# GOARCH=arm
# OUTPATH=./build/linuxarm/
# mkdir -p $OUTPATH
# CGO_ENABLED=0 GOOS="linux" GOARCH="$GOARCH" GOARM=7 go build -ldflags="$LDFLAGS" -o $OUT
# mv $OUT $OUTPATH
# cp ./*.toml $OUTPATH
