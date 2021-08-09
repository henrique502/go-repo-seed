# Directory to place `go install`ed binaries into.
export GOBIN ?= $(shell pwd)/bin

GOLINT = $(GOBIN)/golint

GO_FILES ?= $(shell find . '(' -path .git -o -path vendor ')' -prune -o -name '*.go' -print)

# Also update ignore section in .codecov.yml.
COVER_IGNORE_PKGS =

.PHONY: build
build:
	go build ./...

.PHONY: test
test:
	go test -race ./...

.PHONY: gofmt
gofmt:
	$(eval FMT_LOG := $(shell mktemp -t gofmt.XXXXX))
	gofmt -e -s -l $(GO_FILES) > $(FMT_LOG) || true
	@[ ! -s "$(FMT_LOG)" ] || (echo "gofmt failed:" && cat $(FMT_LOG) && false)

$(GOLINT):
	cd tools && go install golang.org/x/lint/golint

.PHONY: golint
golint: $(GOLINT)
	$(GOLINT) ./...

.PHONY: lint
lint: gofmt golint

# comma separated list of packages to consider for code coverage.
COVER_PKG = $(shell \
	go list -find ./... | \
	grep -v $(foreach pkg,$(COVER_IGNORE_PKGS),-e "^$(pkg)$$") | \
	paste -sd, -)

.PHONY: cover
cover:
	go test -coverprofile=cover.out -coverpkg  $(COVER_PKG) -v ./...
	go tool cover -html=cover.out -o cover.html
