# Web-Service & CLI for Converting SIF to JSON & Vice Versa

## Getting Started

1. Create Web-Service executable.

    Goto `/Server`, run `build.sh "sif-spec(txt) path"`, create executable and its dependencies.

    Goto `./build/your-os/`, make sure 'config.toml' is in same directory.

    Make sure 'config.toml' has correct settings, especially [Cfg2JSON], [Cfg2SIF] and all [File].    

    In [Cfg2JSON], make sure [SIFCfgDir4LIST], [SIFCfgDir4NUM], [SIFCfgDir4BOOL] are correct.

    In [Cfg2SIF], make sure [SIFSpecDir], [ReplCfgPath] are correct.
    
2. Create CLI executable.

    Goto `/Client`, run 'build.sh', create executable.

    Goto `./build/your-os/`, make sure 'config.toml' is in same directory.

3.  Run Web-Service executable with correct config.

4.  Fetch Client executable from `wget` once Web-Service is running. 

    Client usage: e.g. ruled by SIF 3.4.6, want to get JSON from 'Activity.xml' SIF file.

    run `./client SIF2JSON -i=../data/examples/Activity.xml -v=3.4.6`

## Prerequisites

SIF Specification Description File. Text readable format, and at least contains:

   1. Spec VERSION.

   2. LIST, NUMERIC, BOOLEAN attribute type description.
   
   3. Element TRAVERSE description.

   4. `serverEcho.go` must be in `github.com/opentracing-contrib/go-stdlib/nethttp` if you want tracing is available.
      a copy of `serverEcho.go` exists in `/X` folder.
