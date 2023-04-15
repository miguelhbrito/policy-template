SOURCE_FILES := $(shell find . -type f -name '*.go')
# It's necessary to call cut because kwctl command does not handle version
# starting with v.
VERSION ?= $(shell git describe --tags --always | cut -c2-)


policy.wasm: $(SOURCE_FILES) go.mod go.sum types_easyjson.go
	docker run \
		--rm \
		-e GOFLAGS="-buildvcs=false" \
		-v ${PWD}:/src \
		-w /src tinygo/tinygo:0.27.0 \
		tinygo build -o policy.wasm -target=wasi -no-debug .

.PHONY: generate-easyjson
types_easyjson.go: types.go
	docker run \
		--rm \
		-v ${PWD}:/src \
		-w /src \
		golang:1.20-alpine ./hack/generate-easyjson.sh

.PHONY: test
test: types_easyjson.go
	go test -v

.PHONY: lint
lint:
	go vet ./...
	golangci-lint run

.PHONY: clean
clean:
	go clean
	rm -f policy.wasm annotated-policy.wasm artifacthub-pkg.yml