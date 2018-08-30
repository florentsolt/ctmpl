DOCKER					:= docker
DOCKER_IMAGE			:= gotmpl:latest
GOGET 					:= go get -u
GOINSTALL   			:= go install
GOCLEAN   				:= go clean
GOFMT 					:= go fmt
GOLANGCLILINT			:= golangci-lint run
GOLINT					:= golint
GOVET       			:= go vet
GODIRS					:= ./...
GOBENCH     			:= gotest -bench=.
GODOC 					:= godoc -templates ./doc
SED_TOC_LINKS 			:= sed -i -e 's/\]\(.*\)/\]\L\1/'
SED_EMPTY_LINES_CODE 	:= sed -i -e 's/\([[:alnum:]]\)```$$/\1\n```/'

ifdef TEST
GOTEST      := gotest -run ${TEST}
else
GOTEST      := gotest
endif

.PHONY: all fmt lint vet install doc $(GODIRS) docker-build docker-run 

all: fmt vet lint install

docker-build:
	@$(DOCKER) build . -t $(DOCKER_IMAGE)

docker-run:
	@$(DOCKER) run -it -v $$PWD:/go/src/github.com/florentsolt/gotmpl:delegated gotmpl:latest 

install:
	@$(GOINSTALL) ./...

fmt:
	@$(GOFMT) $(GODIRS)

lint:
	@$(GOLANGCLILINT) $(GODIRS)

vet:
	@$(GOVET) $(GODIRS)

test: install
	@find template/testdata -name *.html.go -delete
	@cd template/testdata/benchmark/hero && hero -pkgname hero &> /dev/null
	@cd template/testdata/benchmark/gotmpl && gotmpl &> /dev/null
	@$(GOTEST) $(GODIRS)

bench: install
	@find template/testdata -name *.html.go -delete
	@cd template/testdata/benchmark/hero && hero -pkgname hero &> /dev/null
	@cd template/testdata/benchmark/gotmpl && gotmpl &> /dev/null
	@$(GOBENCH) $(GODIRS)

doc:
	@for i in $(shell go list $(GODIRS)); do \
		$(GOLINT) $$i; \
		$(GODOC) $$i > $$GOPATH/src/$$i/README.md; \
		$(SED_TOC_LINKS) $$GOPATH/src/$$i/README.md; \
		$(SED_EMPTY_LINES_CODE) $$GOPATH/src/$$i/README.md; \
	done

