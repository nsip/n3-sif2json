# Web-Service for Converting SIF to JSON & Vice Versa

## Getting Started

1. Make sure `GOLANG` dev package & `GIT` are available in your machine.

2. Build.
  
   Run `./build.sh` to build service with SIF 3.4.6 & 3.4.7

3. Release.

   Run `release.sh 'dest-platform' 'dest-path'`.

   e.g. run `./release.sh [linux64|win64|mac] ~/Desktop/sif2json/linux64/`
  
4. Docker Deploy (local, only for linux64 platform server).

   Make sure `Docker` is installed.

   Jump into your release dest-path in above step.

   e.g. jump into `~/Desktop/sif2json/linux64/`

   Run `docker build --tag n3-sif2json .` to make docker image.

   Run `docker run --name sif2json --net host n3-sif2json:latest` to run docker image.

5. Test.

   Make sure `curl` is installed.

   curl test script in "test.sh", which goes through all examples in `./data/`.

   Before running `./test.sh`, modify some params like URL (IP, PORT ...) if needed (especially service version number in URL).

   Refer to 'test.sh', write more your own curl test.

## Prerequisites

SIF Specification Description File (txt file). It's text readable format, and at least contains:

1. Spec VERSION.

2. LIST, NUMERIC, BOOLEAN attribute type description.
  
3. Element TRAVERSE description.
