FROM golang:1.14

WORKDIR /go/src/odrv
COPY odrv.go .
COPY go.mod .
RUN go get
RUN go build

ENV PORT=80

EXPOSE 80/tcp
EXPOSE 80/udp

ENTRYPOINT ["/go/src/odrv/odrv"]
