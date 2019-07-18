FROM golang:1.12.5

ADD . /go/src/github.com/tanus-co/snowflake

RUN cd /go/src/github.com/tanus-co/snowflake; go mod download
RUN go install /go/src/github.com/tanus-co/snowflake

ENTRYPOINT /go/bin/snowflake

EXPOSE 22122