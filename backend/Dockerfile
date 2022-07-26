FROM golang:alpine3.15

COPY shortenerBE/ /usr/bin/shortenerBE/

WORKDIR /usr/bin/shortenerBE

RUN go mod tidy

# RUN go build -o /main

ENTRYPOINT ["/bin/sh", "-c", "--"]