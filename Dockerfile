ARG EOSIO_TAG="lishi-zsw-2.0.8-release-v1"
ARG DEB_PKG="lishi-zswchain_2.0.8-dm.12.0_amd64.deb"

FROM ubuntu:18.04 AS base
ARG EOSIO_TAG
ARG DEB_PKG
RUN apt update && apt-get -y install curl ca-certificates libicu60 libusb-1.0-0 libcurl3-gnutls
RUN mkdir -p /var/cache/apt/archives/
RUN curl -sL -o/var/cache/apt/archives/eosio.deb "https://github.com/invisible-train-40/zswchain-lishi/releases/download/${EOSIO_TAG}/${DEB_PKG}"
RUN dpkg -i /var/cache/apt/archives/eosio.deb
RUN rm -rf /var/cache/apt/*

FROM node:12 AS zsw-lishi-launcher
WORKDIR /work
ADD go.mod /work
RUN apt update && apt-get -y install git
RUN cd /work && echo "中数文" && git clone https://github.com/invisible-train-40/zsw-lishi-launcher.git zsw-lishi-launcher &&\
	grep -w github.com/invisible-train-40/zsw-lishi-launcher go.mod | sed 's/.*-\([a-f0-9]*$\)/\1/' |head -n 1 > zsw-lishi-launcher.hash &&\
    cd zsw-lishi-launcher &&\
    git checkout "$(cat ../zsw-lishi-launcher.hash)" &&\
    cd dashboard/client &&\
    yarn install && yarn build

FROM node:12 AS eosq
ADD eosq /work
WORKDIR /work
RUN yarn install && yarn build

FROM golang:1.14 as dfuse
RUN go get -u github.com/GeertJohan/go.rice/rice && export PATH=$PATH:$HOME/bin:/work/go/bin
RUN mkdir -p /work/build
ADD . /work
WORKDIR /work
COPY --from=eosq      /work/ /work/eosq
# The copy needs to be one level higher than work, the dashboard generates expects this file layout
COPY --from=zsw-lishi-launcher /work/zsw-lishi-launcher /zsw-lishi-launcher
RUN cd /zsw-lishi-launcher/dashboard && go generate
RUN cd /work/eosq/app/eosq && go generate
RUN cd /work/dashboard && go generate
RUN cd /work/dgraphql && go generate
RUN go test ./...
RUN go build -v -o /work/build/dfuseeos ./cmd/dfuseeos

FROM base
RUN mkdir -p /app/ && curl -Lo /app/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.2.2/grpc_health_probe-linux-amd64 && chmod +x /app/grpc_health_probe
COPY --from=dfuse /work/build/dfuseeos /app/dfuseeos
COPY --from=dfuse /work/tools/manageos/motd /etc/motd
COPY --from=dfuse /work/tools/manageos/scripts /usr/local/bin/
RUN echo cat /etc/motd >> /root/.bashrc
