FROM alpine:3.6

RUN apk update && \
    apk add  --no-cache ca-certificates bash curl && \
\
    apk add librdkafka --update-cache --repository http://nl.alpinelinux.org/alpine/edge/community && \
\
    apk add musl libressl2.5-libcrypto libressl2.5-libssl && \
\
    rm -rf /var/cache/apk/*
