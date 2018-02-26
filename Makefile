GO_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")
REPOPATH := github.com/r2d4/joust
.PHONY: build
build: $(`GO_FILES)
	CGO_ENABLED=0 go install $(REPOPATH)
	joust

