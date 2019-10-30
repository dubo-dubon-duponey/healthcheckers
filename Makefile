.PHONY: default
default: all

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: build
build:
	go build -v -ldflags "-s -w" -o dist/http-health ./cmd/http/main.go
	go build -v -ldflags "-s -w" -o dist/rtsp-health ./cmd/rtsp/main.go
	go build -v -ldflags "-s -w" -o dist/dns-health ./cmd/dns/main.go

.PHONY: all
all: fmt vet build
