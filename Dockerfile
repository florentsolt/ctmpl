FROM golang:1.10-alpine

RUN apk --no-cache add --update bash git build-base 
RUN go get \
    golang.org/x/tools/cmd/goimports \
    github.com/golangci/golangci-lint/cmd/golangci-lint \
    github.com/rakyll/gotest \
    golang.org/x/net/html

EXPOSE 80
WORKDIR /go/src/github.com/florentsolt/gotmpl

RUN echo "echo -ne '\033]0;Gotmpl\007'" >> /etc/profile
RUN echo "export PATH=$GOPATH/bin:$PATH" >> /etc/profile
