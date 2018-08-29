GOPATH 		:= /go
GOGET 		:= go get -u
GOINSTALL   := go install
GOCLEAN     := go clean
GOFMT 		:= go fmt
GOLINT		:= golangci-lint run
GOVET       := go vet
GODIRS		:= $(shell go list ./...)
GODEPS		:= $(shell go list -f '{{ join .Deps "\n" }}' ./... | sort | uniq | grep "github.com/florentsolt")
GOBENCH     := gotest -bench=.

ifdef TEST
GOTEST      := gotest -run ${TEST}
else
GOTEST      := gotest
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

test: install
	@find template/testdata -name *.html.go -delete
	@cd template/testdata/benchmark/hero && hero -pkgname hero &> /dev/null
	@cd template/testdata/benchmark/gotmpl && gotmpl &> /dev/null
	@$(GOTEST) $(GODIRS) $(GODEPS)

bench: install
	@find template/testdata -name *.html.go -delete
	@cd template/testdata/benchmark/hero && hero -pkgname hero &> /dev/null
	@cd template/testdata/benchmark/gotmpl && gotmpl &> /dev/null
	@$(GOBENCH) $(GODIRS) $(GODEPS)

include $(wildcard *.mk)