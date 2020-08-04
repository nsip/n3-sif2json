# Web-Service & CLI for Converting SIF to JSON & Vice Versa

## Getting Started

0. If no go.mod, Run `go mod init github.com/nsip/n3-sif2json`
  
    If no go.sum, Run `go get -u ./...`

1. Create SIF JSON Configure, Server (Web-Service) and Client (CLI) executables.
  
    Run `build.sh "sif-spec(txt) path"`.  
  
    e.g. run `./build.sh ./SIFSpec/3.4.6.txt ./SIFSpec/3.4.7.txt` to build a web service with SIF 3.4.6 & 3.4.7 AND its CLI Client.

    SIF Config is under ./2JSON/SpecCfg/(version)

    Server executable is under ./Server/build/your-os/

    Client executable is under ./Client/build/your-os/

2. Run Server (Web-Service) executable.

    Goto `./Server/build/your-os/`, make sure 'config.toml' is in this directory.

    Make sure 'config.toml' has correct settings, especially [Cfg2JSON], [Cfg2SIF] and all [File].

    In [Cfg2JSON], make sure [SIFCfgDir4LIST], [SIFCfgDir4NUM], [SIFCfgDir4BOOL] are correct.

    In [Cfg2SIF], make sure [SIFSpecDir], [ReplCfgPath] are correct.
  
3. Check Client (CLI) executable (optional).

    Goto `./Client/build/your-os/`, make sure 'config.toml' is in this directory.

4. Fetch Client executable and its configure from `wget` when Web-Service is running. 

    e.g. `wget ip:port/client-linux64`, `wget -O config.toml ip:port/client-config`

    Client Usage: e.g. for SIF-3.4.6, get JSON from 'Activity.xml' SIF file.

    Run `./client SIF2JSON -i=../data/examples/Activity.xml -v=3.4.6`

## Prerequisites

SIF Specification Description File. Text readable format, and at least contains:

1. Spec VERSION.

2. LIST, NUMERIC, BOOLEAN attribute type description.
  
3. Element TRAVERSE description.

## Deployment

1. Copy `Dockerfile` to ../

2. Run `docker build --tag=n3-sif2json .`

## Others

1. Make sure /Server/config.toml [Cfg2JSON] [Cfg2SIF] are correct.

2. UnitTest Mode, Set /2JSON/config.toml [SIFCfgDir4LIST], [SIFCfgDir4NUM], [SIFCfgDir4BOOL] to `../`;

3. UnitTest Mode, Set /2SIF/config.toml [SIFSpecDir], [ReplCfgPath] to `../`.

4. Server Mode, Set /2JSON/config.toml [SIFCfgDir4LIST], [SIFCfgDir4NUM], [SIFCfgDir4BOOL] to `../../../`;

5. Server Mode, Set /2SIF/config.toml [SIFSpecDir], [ReplCfgPath] to `../../../`.

6. Server Mode, Make sure /Preprocess/genStruct_test.go `flag.Args()` is working.

7. Using goclient, Make sure its config.toml [Route] is correct.
