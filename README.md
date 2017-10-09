Alpine images for build/run confluent kafka clients
===================================================

Confluent kafka go client lib is very fast because it uses C/C++ librdkafka lib, which also means go clients built this way cannot be completely static linked (need glibc). For minial docker images, it is often recommended to use alpine/muslcc instead gcc/glibc.

Dockerfile.dev_kafka_go:

        dev env to build go clients.
        get it: 'docker pull yglcode/alpine-kafka-go-dev'
        
Dockerfile.run_kafka_go:

        runtime env to run go client binaries built above.
        get it: 'docker pull yglcode/alpine-kafka-go-run'

In a multi-stage Dockerfile, we can use Dockerfile.dev_kafka_go at early stage and Dockerfile.run_kafka_go as last stage, to produce a minimal deploy image, as shown in sample "pingpong".

