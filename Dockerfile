FROM alpine
RUN mkdir /n3-sif2json
COPY ./n3-sif2json /n3-sif2json
WORKDIR /n3-sif2json/Server/build/linux64
CMD ["./server"]

### ! run this Dockerfile at parent dir of n3-sif2json 
### docker build --tag=n3-sif2json . 

### docker tag IMAGE_ID cdutwhu/n3-sif2json:latest
### docker login
### docker push cdutwhu/n3-sif2json