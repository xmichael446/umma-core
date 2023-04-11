# docker build . -t cosmoscontracts/umma:latest
# docker run --rm -it cosmoscontracts/umma:latest /bin/sh
FROM golang:1.18-alpine3.15 AS go-builder

# this comes from standard alpine nightly file
#  https://github.com/rust-lang/docker-rust-nightly/blob/master/alpine3.12/Dockerfile
# with some changes to support our toolchain, etc
SHELL ["/bin/ash", "-eo", "pipefail", "-c"]
# we probably want to default to latest and error
# since this is predominantly for dev use
# hadolint ignore=DL3018
RUN set -eux; apk add --no-cache ca-certificates build-base;

# hadolint ignore=DL3018
RUN apk add git
# NOTE: add these to run with LEDGER_ENABLED=true
# RUN apk add libusb-dev linux-headers

WORKDIR /code
COPY . /code/



# OLD VERSION v1.1.1
# See https://github.com/CosmWasm/wasmvm/releases
#ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.1.1/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
#ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.1.1/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a

#VERSION v1.1.1
COPY /cache_libs/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
COPY /cache_libs/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a

#VERSION v1.1.1
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 9ecb037336bd56076573dc18c26631a9d2099a7f2b40dc04b6cae31ffb4c8f9a
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep 6e4de7ba9bad4ae9679c7f9ecf7e283dd0160e71567c6a7be6ae47c81ebe7f32


# NEW VERSION
#ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.2.2/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
#ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.2.2/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a

# NEW VERSION
#COPY /cache_libs_v12/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
#COPY /cache_libs_v12/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a

# NEW VERSION
#RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 20ab42fceddd5347b973c254717eed62b2d46925c098f58304d09488ed59464a
#RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep fc859817a1c548ceb0e02ad8cf5c8e374d240caac547d04dd02b284ace8ab33d

# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp "/lib/libwasmvm_muslc.$(uname -m).a" /lib/libwasmvm_muslc.a

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
# then log output of file /code/bin/ummad
# then ensure static linking
RUN LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true make build \
  && file /code/bin/ummad \
  && echo "Ensuring binary is statically linked ..." \
  && (file /code/bin/ummad | grep "statically linked")

# --------------------------------------------------------
FROM alpine:3.15

COPY --from=go-builder /code/bin/ummad /usr/bin/ummad
#RUN apt update && apt install -y python3 protobuf-compiler

COPY docker/* /opt/
RUN chmod +x /opt/*.sh

WORKDIR /opt

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657

CMD ["/usr/bin/ummad", "version"]