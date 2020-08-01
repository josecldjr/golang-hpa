# BUILDER IMAGE
FROM golang:alpine3.11
WORKDIR /go/src/app

# compile go scripts
COPY . /go/src/app

RUN export GOPATH=/go/src/
RUN export GOROOT=/go/src/app
RUN export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

RUN apk add git
RUN apk add gcc

RUN go get /go/src/app

RUN chmod +x .
# RUN go install 
