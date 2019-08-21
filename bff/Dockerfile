FROM golang:1.12 AS bff-builder

RUN apt-get update && \
  apt-get install -y git unzip && \
  apt-get clean && rm -rf /var/lib/apt/lists/*

ENV GO111MODULE on

WORKDIR /go/src/github.com/seriwb/front-bff-msa/bff
COPY . .
RUN make


FROM alpine AS bff-server

EXPOSE 8000
COPY --from=bff-builder /go/src/github.com/seriwb/front-bff-msa/bff/build/bff_server .