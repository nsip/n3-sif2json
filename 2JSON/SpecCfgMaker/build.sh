 #!/bin/bash

### Create JSON files at building time ###

go get

CGO_ENABLED=0 go run var.go spec.go maker.go -- ./List2JSON.toml ./Num2JSON.toml ./Bool2JSON.toml