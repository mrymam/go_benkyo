FROM golang:1.17

WORKDIR /go/src/app
COPY go.mod ./
COPY go.sum ./

RUN go mod download \
  && go get -u github.com/cosmtrek/air@v1.29.0 \
  && go install github.com/go-delve/delve/cmd/dlv@latest