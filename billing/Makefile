APPNAME=billing

.DEFAULT_GOAL := help

## test: run tests on cmd and pkg files.
.PHONY: test
test: vet fmt
	go test ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

## build: build application binary.
.PHONY: build
build:
	go build -o $(APPNAME)

## run: run the api
.PHONY: run
run:
	go run .

.PHONY: install_goose
install_goose:
	$(call install_goose)

## install-deps: install main dependencies
install-deps:
	@echo "installing golangci-lint..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.23.8
	@echo "golangci installed!"
	@echo "installing errcheck..."
	go install github.com/kisielk/errcheck@latest
	@echo "errcheck installed!"
	@echo "installing goose..."
	go install github.com/pressly/goose/v3/cmd/goose@latest
	@echo "goose installed!"
	@echo "installing swagger..."
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	@echo "swagger installed!"


check-lint:
ifeq (, $(shell which golangci-lint))
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.23.8
endif
ifeq (, $(shell which errcheck))
	go install github.com/kisielk/errcheck@latest
endif

## lint: run linters over the entire code base
.PHONY: lint
lint: check-lint
	golangci-lint run ./... --timeout 15m0s
	errcheck -exclude ./.golangci-errcheck-exclude.txt ./...

## install-hooks: install hooks
.PHONY: install-hooks
install-hooks:
	ln -s $(PWD)/githooks/pre-push ../.git/hooks/pre-push

## clean: remove releases
.PHONY: clean
clean:
	rm -rf $(APPNAME)

all: help
.PHONY: help
help: Makefile
	@echo " Choose a command..."
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
