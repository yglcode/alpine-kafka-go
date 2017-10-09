FROM yglcode/alpine-kafka-go-dev as builder
LABEL builder=true
ENV srcDir="/go/src/github.com/yglcode/alpine-kafka-go/samples/pingpong"
RUN mkdir -p $srcDir
COPY . $srcDir/
WORKDIR $srcDir
RUN go get ./... && go build -o pingpong

FROM yglcode/alpine-kafka-go-run
ENV srcDir="/go/src/github.com/yglcode/alpine-kafka-go/samples/pingpong"
COPY --from=builder $srcDir/pingpong .
CMD ["/pingpong"]
