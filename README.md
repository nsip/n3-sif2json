# Web-Service & CLI for Converting SIF to JSON & Vice Versa

## Getting Started

1. Create Web-Service executable.

    Goto `/Server`, run `build.sh "sif-spec(txt) path"`, create executable and its dependencies.

    e.g. run `./build.sh ../SIFSpec/3.4.6.txt` to build a web service with SIF 3.4.6 for different OS.

    Goto `./build/your-os/`, make sure 'config.toml' is in same directory.

    Make sure 'config.toml' has correct settings, especially [Cfg2JSON], [Cfg2SIF] and all [File].    

    In [Cfg2JSON], make sure [SIFCfgDir4LIST], [SIFCfgDir4NUM], [SIFCfgDir4BOOL] are correct.

    In [Cfg2SIF], make sure [SIFSpecDir], [ReplCfgPath] are correct.
    
2. Create CLI executable.

    Goto `/Client`, run 'build.sh', create executable.

    Goto `./build/your-os/`, make sure 'config.toml' is in same directory.

3.  Run Web-Service executable with correct config.

4.  Fetch Client executable and its configure from `wget` when Web-Service is running. 

    e.g. `wget ip:port/client-linux64`, `wget -O config.toml ip:port/client-config`

    Client Usage: e.g. for SIF-3.4.6, get JSON from 'Activity.xml' SIF file.

    Run `./client SIF2JSON -i=../data/examples/Activity.xml -v=3.4.6`

## Prerequisites

SIF Specification Description File. Text readable format, and at least contains:

   1. Spec VERSION.

   2. LIST, NUMERIC, BOOLEAN attribute type description.
   
   3. Element TRAVERSE description.

   4. `serverEcho.go` must be in `github.com/opentracing-contrib/go-stdlib/nethttp` if you want tracing is available.
      a copy of `serverEcho.go` exists in `/X` folder.
