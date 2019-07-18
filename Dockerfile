FROM golang:1.12.5

ADD ./snowflake /bin/
ADD ./app.yml /go/

ENTRYPOINT /bin/snowflake

EXPOSE 50051