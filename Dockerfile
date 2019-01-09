FROM golang:1.11-stretch AS build-env

COPY . /ropsten-faucet
WORKDIR /ropsten-faucet
RUN go build -v -mod=vendor -o /cmd

FROM debian:stretch
COPY --from=build-env /cmd /

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/cmd"]
