# create image from the official Go image
FROM golang:alpine
RUN apk add --update tzdata \
    bash wget curl git;
# Create binary directory, install fresh
RUN mkdir -p $$GOPATH/bin && \
    go get github.com/pilu/fresh
# define work directory
ADD . /go/src/sample-service2
WORKDIR /go/src/sample-service2
# serve the app
EXPOSE 4000
CMD go get && fresh -c runner.conf main.go