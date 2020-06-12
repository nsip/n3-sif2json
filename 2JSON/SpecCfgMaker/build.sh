 #!/bin/bash

### Create JSON files at building time ###

go get

go run var.go spec.go maker.go -- ./List2JSON.toml ./Num2JSON.toml ./Bool2JSON.toml