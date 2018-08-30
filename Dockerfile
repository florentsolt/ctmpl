FROM golang:1.11-alpine

RUN apk --no-cache add --update git build-base 
RUN go get \
    golang.org/x/tools/cmd/godoc \
    golang.org/x/lint/golint \
    github.com/golangci/golangci-lint/cmd/golangci-lint \
    github.com/rakyll/gotest \
    golang.org/x/net/html \
    golang.org/x/tools/cmd/goimports \
    github.com/shiyanhui/hero/hero \
    github.com/dchest/htmlmin

WORKDIR /go/src/github.com/florentsolt/gotmpl
