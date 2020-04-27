# Web-Service & CLI for Converting SIF to JSON & Vice Versa

## Getting Started

1. Convert SIF Spec Description (txt) to toml files.

    ( !!! IGNORE THIS STEP unless 3 folders, [BOOLEAN, LIST, NUMERIC], under /2JSON/SpecCfg/(sif-version)/ DO NOT exist !!! )

    goto /SpecProcess, run 'build.sh', create executable.

    run executable with 'SIF Specification Description File',
                        'go source base file (for config.go)',
                        'List2JSON base file (for List2JSON.toml)',
                        'Num2JSON base file (for Num2JSON.toml)',
                        'Bool2JSON base file (for Bool2JSON.toml)'.

    to generate 'config.go',
                'Bool2JSON.toml',
                'List2JSON.toml',
                'Num2JSON.toml'.

    e.g. generate SIF 3.4.6
    run `./SpecProcess ../SIFSpec/3.4.6.txt ../2JSON/SpecCfgMaker/base-go/config ../2JSON/SpecCfgMaker/base-toml/List2JSON ../2JSON/SpecCfgMaker/base-toml/Num2JSON ../2JSON/SpecCfgMaker/base-toml/Bool2JSON ../2JSON/SpecCfgMaker/`

2. Create Configure(json) from toml files.

    ( !!! IGNORE THIS STEP unless 3 folders, [BOOLEAN, LIST, NUMERIC], under /2JSON/SpecCfg/(sif-version)/ DO NOT exist !!! )

    goto /2JSON/SpecCfgMaker, run 'build.sh'
                              with 'config.go' from step 1, create executable.

    run executable with 'List2JSON.toml',
                        'Num2JSON.toml'
                        'Bool2JSON.toml' from step 1,

    to generate Spec Configure(json)

    e.g. generate SIF 3.4.6
    run `./mkSpecCfg ./List2JSON.toml ./Num2JSON.toml ./Bool2JSON.toml`

3. Create Web-Service executable.

    goto /Server, run 'build.sh', create executable.

    run executable with 'config.toml' which has Port setting etc.
    (a copy of config.toml exists in /Server/config)

    make SIF 3.4.6 service running. (default is Linux executable, goto build/your-os to run your OS)
    run `./server`

4. Create CLI executable.

    goto /Client, run 'build.sh', create executable.

    run executable with 'config.toml' which has IP setting etc.
    (a copy of config.toml exists in /Server/config)

    e.g. ruled by SIF 3.4.6, get JSON from 'Activity.xml' SIF file. (default is Linux executable, goto build/your-os to run your OS)
    run `./client SIF2JSON -i=../data/examples/Activity.xml -v=3.4.6`

## Prerequisites

SIF Specification Description File.
Text readable format, and at least contains:

   1. Spec VERSION.
   2. LIST, NUMERIC, BOOLEAN attribute type description.
   3. Element TRAVERSE description.
