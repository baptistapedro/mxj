FROM ubuntu:20.04 as builder

RUN ln -snf /usr/share/zoneinfo/$CONTAINER_TIMEZONE /etc/localtime && echo $CONTAINER_TIMEZONE > /etc/timezone

RUN DEBIAN_FRONTEND=noninteractive \
	apt-get update && apt-get install -y build-essential tzdata pkg-config \
	wget clang git

RUN wget https://go.dev/dl/go1.19.1.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.1.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

ADD . /mxj
WORKDIR /mxj
ADD fuzzers/fuzz_newmapxml.go ./fuzzers/
WORKDIR ./fuzzers/
RUN go mod init myfuzz
RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/clbanning/mxj
RUN /root/go/bin/go-fuzz-build -libfuzzer -o harness.a
RUN clang -fsanitize=fuzzer harness.a -o fuzz_newmapxml

FROM ubuntu:20.04
COPY --from=builder /mxj/fuzzers/fuzz_newmapxml /
COPY --from=builder /mxj/*.xml /testsuite/

ENTRYPOINT []
CMD ["/fuzz_newmapxml"]