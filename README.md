# Web-Service for Converting SIF to JSON & Vice Versa

## Getting Started

0. Run `go get -u ./...` to update this project's dependencies.
   Ignore any `undefined: n3cfg.***` errors.

1. Build.

   In 'build.sh', change 'password' in line "sudopwd='password'" to your real sudo/admin password & save.
  
   Run `build.sh 'paths of sif-spec(txt)'`.  
  
   e.g. run `./build.sh ./SIFSpec/3.4.6.txt ./SIFSpec/3.4.7.txt` to build services with SIF 3.4.6 & 3.4.7

2. Release.

   Run `release.sh 'dest-platform' 'dest-path'`.

   e.g. run `./release.sh [linux64|win64|mac] ~/Desktop/sif2json/linux64/`
  
3. Docker Deploy (local, only for linux64 platform server).

   Make sure Docker installed.

   Jump into your release dest-path in above step 2.

   e.g. jump into `~/Desktop/sif2json/linux64/`

   Run `docker build --tag=n3-sif2json .` to make docker image.

   Run `docker run --name sif2json --net host n3-sif2json:latest` to run docker image.

4. Test.

   Simple curl test scripts in test.sh.

   Before running `./test.sh`, modify some params like URL (IP, PORT ...) if needed.

   OR write your own curl test like 'test.sh'.

## Prerequisites

SIF Specification Description File. It's text readable format, and at least contains:

1. Spec VERSION.

2. LIST, NUMERIC, BOOLEAN attribute type description.
  
3. Element TRAVERSE description.

## Remind for who plays with this source

1. Make sure /config.toml [Cfg2JSON] [Cfg2SIF] are correct.

2. For UnitTest, Set /2JSON/config.toml [SIFCfgDir4LIST], [SIFCfgDir4NUM], [SIFCfgDir4BOOL] to `../`;

3. For UnitTest, Set /2SIF/config.toml [SIFSpecDir], [ReplCfgPath] to `../`.

4. For Server, Set /2JSON/config.toml [SIFCfgDir4LIST], [SIFCfgDir4NUM], [SIFCfgDir4BOOL] to `../../../`;

5. For Server, Set /2SIF/config.toml [SIFSpecDir], [ReplCfgPath] to `../../../`.
