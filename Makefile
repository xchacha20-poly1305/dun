NAME = dun
COMMIT = $(shell git rev-parse --short HEAD)
TAGS = with_conntrack,with_gvisor,with_wireguard,with_clash_api,with_utls,with_ech
TAGS_TEST ?= with_gvisor,with_quic,with_wireguard,with_grpc,with_ech,with_utls,with_reality_server

GOHOSTOS = $(shell go env GOHOSTOS)
GOHOSTARCH = $(shell go env GOHOSTARCH)
# VERSION=$(shell CGO_ENABLED=0 GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) go run ./cmd/internal/read_tag)

PARAMS = -v -trimpath -ldflags "-X 'github.com/sagernet/sing-box/constant.Version=$(VERSION)' -s -w -buildid="
MAIN_PARAMS = $(PARAMS) -tags $(TAGS)
MAIN = ./
PREFIX ?= $(shell go env GOPATH)

build:
	go build $(MAIN_PARAMS) $(MAIN)

ci_build:
	go build $(PARAMS) $(MAIN)
	go build $(MAIN_PARAMS) $(MAIN)

install:
	go build -o $(PREFIX)/bin/$(NAME) $(MAIN_PARAMS) $(MAIN)

fmt:
	@gofumpt -l -w .
	@gofmt -s -w .
	@gci write --custom-order -s standard -s "prefix(github.com/sagernet/)" -s "default" .

fmt_install:
	go install -v mvdan.cc/gofumpt@latest
	go install -v github.com/daixiang0/gci@latest

lint:
	GOOS=linux golangci-lint run ./...
	GOOS=android golangci-lint run ./...
	GOOS=windows golangci-lint run ./...
	GOOS=darwin golangci-lint run ./...
	GOOS=freebsd golangci-lint run ./...

lint_install:
	go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest

test:
	@go test -v ./... && \
	cd test && \
	go mod tidy && \
	go test -v -tags "$(TAGS_TEST)" .

test_stdio:
	@go test -v ./... && \
	cd test && \
	go mod tidy && \
	go test -v -tags "$(TAGS_TEST),force_stdio" .

update:
	git fetch
	git reset FETCH_HEAD --hard
	git clean -fdx