#
# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
#
# BUILD: docker build -t nsip/sif2json .
# RUN: docker run -it -p1324:1324 nsip/sif2json
############################
# STEP 1 build executable binary
############################
# TODO - redo the build in the docker image, for now, assume it is built
#FROM golang:1.13-stretch as builder
#COPY . .
#RUN go get
#RUN go get github.com/labstack/echo/middleware
##RUN go get github.com/nsip/sifclient-go
##RUN go get github.com/nsip/sifdata-go
#RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/hitsproxy cmd/hitsproxy/main.go
############################
# STEP 2 build a small image
############################
FROM debian:stretch
COPY build/Linux64/config.toml /go/bin/config.toml
COPY build/Linux64/server /go/bin/server
WORKDIR /go/bin/
CMD ["/go/bin/server"]
