FROM golang:alpine

MAINTAINER Maintainer

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR .

COPY src /github.com/perajim/src

# Run the two commands below to install git and dependencies for the project. 
# RUN apk update && apk add --no-cache git
# RUN go get ./...

COPY dependencies /go/src #if you don't want to pull dependencies from git 

RUN go build go-docker-dev.to/src/app

EXPOSE $PORT

ENTRYPOINT ["./app"]