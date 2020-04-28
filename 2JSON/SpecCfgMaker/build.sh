 #!/bin/bash

### Create JSON files at building time ###

go run var.go config.go maker.go -- ./List2JSON.toml ./Num2JSON.toml ./Bool2JSON.toml