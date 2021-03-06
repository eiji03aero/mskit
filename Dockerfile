FROM golang:1.14.1-buster
USER root

WORKDIR /app

RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y \
  vim less \
  zip unzip netcat

ENV PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP \
  && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc \
  && unzip -o $PROTOC_ZIP -d /usr/local 'include/*' \
  && rm -f $PROTOC_ZIP

RUN go get -u \
  github.com/golang/protobuf/protoc-gen-go \
  github.com/go-delve/delve/cmd/dlv \
  github.com/cespare/reflex

CMD ["/bin/bash"]
