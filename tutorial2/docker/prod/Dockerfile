FROM golang:1.17 as builder

RUN mkdir /build
COPY . /build/
WORKDIR /build

RUN CGO_ENABLED=0 GO111MODULE=on go build -o server cmd/server.go cmd/router.go

FROM scratch
COPY --from=builder /build/server /go/bin/server
WORKDIR /app
EXPOSE 9000
CMD ["/go/bin/server"]