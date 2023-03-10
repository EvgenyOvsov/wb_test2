NAME=service
VERSION?=$$(git rev-parse HEAD)

default: release

.PHONY: version
version:
	@echo ${VERSION}; exit 0;

.PHONY: build
build: compile compress

.PHONY: compile
compile:
	# The -w turns off DWARF debugging information
	# The -s turns off generation of the Go symbol table
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 \
		go build \
			-mod vendor \
			-ldflags="-w -s -X main.Version=${VERSION}" \
			-o bin/${NAME} \
			./cmd/main.go

.PHONY: compress
compress:
	# check if upx exists
	@if upx --version > /dev/null 2>&1; \
	then \
		upx -9 --ultra-brute bin/${NAME}; \
	else \
	    echo "upx is not installed. Anyway, your executables are compiled."; exit 0; \
	fi

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: run
run:
	go run -mod vendor cmd/main.go

.PHONY: release
release: build