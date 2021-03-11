FROM golang:1.14

WORKDIR /go/src/doh-proxy
COPY main.go .
COPY go.mod .
RUN go get
RUN go build -o main

ENV PORT=80

EXPOSE 80/tcp
EXPOSE 80/udp

ENTRYPOINT ["/go/src/doh-proxy/main"]
