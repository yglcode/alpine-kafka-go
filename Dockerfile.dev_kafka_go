FROM golang:1.9.1-alpine3.6

RUN apk update && \
    apk add --no-cache ca-certificates bash curl git make gcc g++ python zlib cyrus-sasl lz4 libsasl bsd-compat-headers py-setuptools

WORKDIR /src

RUN curl -sSL -O https://github.com/edenhill/librdkafka/archive/master.zip && \
    unzip master.zip && \
    cd /src/librdkafka-master && \
    ./configure && \
    make && \
    make install && \
    rm -fr /src/librdkafka-master /src/master.zip 

RUN curl -sSL -O https://github.com/edenhill/trivup/archive/master.zip && \
    unzip master.zip && \
    cd /src/trivup-master && \
    python setup.py install && \
    rm -fr /src/master.zip /src/trivup-master && \
\
    rm -rf /var/cache/apk/*     

