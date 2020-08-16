FROM alpine
RUN mkdir /n3-sif2json
COPY . / /n3-sif2json/
WORKDIR /n3-sif2json/
CMD ["./server"]

### ! run this Dockerfile 
### docker build --tag=n3-sif2json . 

### ! run this docker image
### docker run --name sif2json --net host n3-sif2json:latest

### ! push image to docker hub
### docker tag IMAGE_ID dockerhub-user/n3-sif2json:latest
### docker login
### docker push dockerhub-user/n3-sif2json