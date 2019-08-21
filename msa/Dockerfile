FROM golang:1.12 AS msa-builder

RUN apt-get update && \
  apt-get install -y git unzip && \
  apt-get clean && rm -rf /var/lib/apt/lists/*

ENV GO111MODULE on
RUN go get -u goa.design/goa/v3/cmd/goa

# install protoc
ARG PROTOC_VER="3.9.1"
RUN mkdir /protoc
ADD https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/protoc-${PROTOC_VER}-linux-x86_64.zip /protoc
WORKDIR /protoc
RUN unzip protoc-${PROTOC_VER}-linux-x86_64.zip && \
  mkdir -p /usr/local/protoc && \
  cp -r bin /usr/local/protoc && \
  chmod 644 -R /usr/local/protoc && \
  chmod 755 /usr/local/protoc/bin/* && \
  cp -r include/google /usr/local/include && \
  chmod 644 -R /usr/local/include/google
ENV PATH $PATH:/usr/local/protoc/bin
WORKDIR /
RUN rm -rf /protoc

RUN go get -u github.com/golang/protobuf/protoc-gen-go

WORKDIR /go/src/github.com/seriwb/front-bff-msa/msa
COPY . .
RUN make


FROM alpine AS msa-server

EXPOSE 8000
COPY --from=msa-builder /go/src/github.com/seriwb/front-bff-msa/msa/build/msa_server .