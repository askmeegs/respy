
FROM alpine
RUN apk add --update ca-certificates && \
      rm -rf /var/cache/apk/* /tmp/*

FROM golang:1.12 as build
COPY . $GOPATH/src/github.com/askmeegs/respy
WORKDIR $GOPATH/src/github.com/askmeegs/respy
RUN go build -o ./respy ./
CMD exec /bin/sh -c "trap : TERM INT; (while true; do sleep 1000; done) & wait"
