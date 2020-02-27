# A Web-Service & Its Command Line Client for Converting SIF to JSON & Vice Versa.

## Getting Started

1. Convert SIF Spec(txt) to toml files.

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

2. Create Configure(json) from toml files.

    goto /2JSON/SpecCfgMaker, run 'build.sh'
                              with 'config.go' from step 1, create executable.

    run executable with 'List2JSON.toml',
                        'Num2JSON.toml'
                        'Bool2JSON.toml' from step 1,

    to generate Spec Configure(json)

3. Create Web-Service executable.

    goto /Server, run 'build.sh', create executable.

    run executable with 'config.toml' which has Port setting etc.
    (a copy of config.toml exists in /Server/config)

4. Create CLI executable.

    goto /Client, run 'build.sh', create executable.

    run executable with 'config.toml' which has IP setting etc.
    (a copy of config.toml exists in /Server/config)

## Prerequisites

SIF Specification Description File.
Text readable format, and at least contains:

   1. Spec VERSION.
   2. LIST, NUMERIC, BOOLEAN attribute type description.
   3. Element TRAVERSE description.
