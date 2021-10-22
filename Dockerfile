FROM debian:bullseye AS libpostal

# apt dependencies
RUN apt-get update \
  && apt-get install -y \
    autoconf \
    automake \
    libtool \
    pkg-config \
    python \
    curl \
    unzip \
    gcc \
    libc6-dev \
    make

# download libpostal src
ENV GIT_HASH 9c975972985b54491e756efd70e416f18ff97958
RUN curl -L \
  https://github.com/openvenues/libpostal/archive/$GIT_HASH.zip \
  > archive.zip
RUN unzip archive.zip -d /src
# RUN ls -l
# RUN ls -l /libpostal-$GIT_HASH

# build libpostal
WORKDIR /src/libpostal-$GIT_HASH
# RUN ls -l
RUN ./bootstrap.sh \
  && ./configure \
    --prefix=/libpostal \
    --datadir=/libpostal/data \
  && make -j4 \
  && make check \
  && make install \
  && ldconfig

FROM golang:1.17 AS builder
COPY --from=libpostal /libpostal /libpostal
COPY . /src
WORKDIR /src

ENV PKG_CONFIG_PATH /libpostal/lib/pkgconfig
ENV LD_LIBRARY_PATH /libpostal/lib
RUN go mod tidy
RUN go generate ./...
RUN go test ./...
RUN go build

FROM debian:bullseye
COPY --from=libpostal /libpostal/lib /libpostal/lib
COPY --from=libpostal /libpostal/data /libpostal/data
COPY --from=builder /src/address-parser /address-parser
ENV LD_LIBRARY_PATH /libpostal/lib
CMD ["/address-parser"]
