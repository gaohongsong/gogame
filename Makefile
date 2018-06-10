
# define build version
VERSION ?= 1.0.0
BUILD_COMMIT := `git rev-parse HEAD`
BUILD_TIME := `date`
GO_VERSION := `go version`

# global go tools
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOTOOL=$(GOCMD) tool
GOFMT=gofmt

# go build flags
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X 'main.BuildCommit=$(BUILD_COMMIT)' \
-X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GoVersion=$(GO_VERSION)'"

# binary name
BINARY := goplayer
PLATFORMS := windows linux darwin
OS = $(word 1, $@)

GO_FILES=$(shell find . -type f -name "*.go" -not -path "./pkg/*")

release: windows linux darwin

# support multi platform
$(PLATFORMS):
	mkdir -p release/$(VERSION)
	GOOS=$(OS) GOARC=amd64 $(GOBUILD) $(LDFLAGS) -o release/$(VERSION)/$(BINARY)-$(VERSION)-$(OS)-amd64

package:
	rm -rf release_$(VERSION).tar.gz
	tar -zcvf release_$(VERSION).tar.gz release
	@echo "create release_$(VERSION).tar.gz successfully"

.PHONY:gotest
gentest:
	gotests -all -excl main -w $(GO_FILES)

.PHONY:test
test:
	rm -rf cover.out
	#$(GOTEST) -v -cover=true -coverprofile=cover.out ./...
	$(GOTEST) -v -cover=true --covermode=count -coverprofile=cover.out ./...

.PHONY:cover
cover:test
	$(GOTOOL) cover -func=cover.out
	#$(GOTOOL) cover -html=cover.out

.PHONY:fmt
fmt:
	$(GOFMT) -l -w $(GO_FILES)

.PHONY:check
check:
	@test -z $(shell gofmt -l $(GO_FILES) | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet ${GO_FILES}

.PHONY:clean
clean:
	$(GOCLEAN)
	rm -rf release
	rm -rf *.out

.PHONY: release $(PLATFORMS)
