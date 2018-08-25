GOPATH 		:= /go
GOGET 		:= go get -u
GOINSTALL   := go install
GOCLEAN     := go clean
GOFMT 		:= go fmt
GOLINT		:= golangci-lint run
GOVET       := go vet
GODIRS		:= $(shell go list ./...)
GODEPS		:= $(shell go list -f '{{ join .Deps "\n" }}' ./... | sort | uniq | grep "github.com/florentsolt")

ifdef TEST
GOTEST      := gotest -p 1 -run ${TEST}
else
GOTEST      := gotest -p 1
endif

.PHONY: all fmt lint vet install

all: fmt install

install:
	@$(GOINSTALL) ./...

fmt:
	@$(GOFMT) $(GODIRS) $(GODEPS)

lint:
	@cd $$GOPATH/src && $(GOLINT) $(GODIRS) $(GODEPS)

vet:
	@$(GOVET) $(GODIRS) $(GODEPS)

test:
	@$(GOTEST) $(GODIRS) $(GODEPS)

include $(wildcard *.mk)