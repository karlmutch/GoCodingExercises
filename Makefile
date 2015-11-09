EXAMPLES := stub

export CGO_ENABLED=0
export GOGC=off

DATE := $(shell date '+%Y-%m-%d_%H:%M:%S%z')
HASH := $(shell git rev-parse HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
GO_VERSION := -ldflags "-X common.buildTime=$(DATE) -X common.gitHash=$(HASH) -X common.gitBranch=$(BRANCH)"

GO_GET := go get -t $(GO_VERSION)
GO_TEST := go test $(GO_VERSION)

.PHONY: all tools modules

all: modules


modules:
	@for example in $(EXAMPLES); do \
          set -e; \
	  pushd src/$$example > /dev/null; \
	  $(GO_TEST) ;\
	  popd > /dev/null; \
	done

tools:
	go get -f -u github.com/golang/lint/golint
	# go extensions, commands
	go get -f -u golang.org/x/tools/cmd/...
	# Go Symbols and SCA package
	go get -f -u github.com/rogpeppe/godef
	# GoCode symbol server
	go get -f -u github.com/nsf/gocode
	# Oracle UI
	go get -f -u github.com/fzipp/pythia
	# Error checking and processing lint
	go get -f -u github.com/kisielk/errcheck
	# Go testing package asserts
	go get github.com/bmizerany/assert
	# Coverage reporting tool
	go get github.com/axw/gocov/gocov


clean:
	@rm -r bin/*
	@rm -r pkg/*
