## docker image prune
## docker rmi $(docker images -a -q)

# FROM alpine
# RUN mkdir /n3-sif2json
# COPY . / /n3-sif2json/
# WORKDIR /n3-sif2json/
# CMD ["./server"]

### ! run this Dockerfile 
### docker build --tag=n3-sif2json . 

### ! run this docker image
### docker run --name sif2json --net host n3-sif2json:latest

### ! push image to docker hub
### docker tag IMAGE_ID dockerhub-user/n3-sif2json:latest
### docker login
### docker push dockerhub-user/n3-sif2json


###########################
# INSTRUCTIONS
############################
# BUILD
#	docker build --rm -t nsip/n3-sif2json:latest -t nsip/n3-sif2json:v0.1.0 .
# TEST: docker run -it -v $PWD/test/data:/data -v $PWD/test/config.json:/config.json nsip/n3-sif2json:develop .
# RUN: docker run -d nsip/n3-sif2json:develop
#
# PUSH
#	Public:
#		docker push nsip/n3-sif2json:v0.1.0
#		docker push nsip/n3-sif2json:latest
#
#	Private:
#		docker tag nsip/n3-sif2json:v0.1.0 the.hub.nsip.edu.au:3500/nsip/n3-sif2json:v0.1.0
#		docker tag nsip/n3-sif2json:latest the.hub.nsip.edu.au:3500/nsip/n3-sif2json:latest
#		docker push the.hub.nsip.edu.au:3500/nsip/n3-sif2json:v0.1.0
#		docker push the.hub.nsip.edu.au:3500/nsip/n3-sif2json:latest
#
###########################
# DOCUMENTATION
############################

###########################
# STEP 0 Get them certificates
############################
# (note, step 2 is using alpine now) 
# FROM alpine:latest as certs

############################
# STEP 1 build executable binary (go.mod version)
############################
FROM golang:1.15.2-alpine3.12 as builder
RUN apk add --no-cache ca-certificates
RUN apk update && apk add --no-cache git bash
RUN mkdir -p /n3-sif2json
COPY . / /n3-sif2json/
WORKDIR /n3-sif2json/
RUN ["/bin/bash", "-c", "./build_d.sh"]
RUN ["/bin/bash", "-c", "./release_d.sh"]

############################
# STEP 2 build a small image
############################
FROM alpine
COPY --from=builder /n3-sif2json/app/ /
# NOTE - make sure it is the last build that still copies the files
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /
CMD ["./server"]
