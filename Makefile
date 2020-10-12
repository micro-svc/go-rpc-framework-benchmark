GOBUILD = go build -ldflags "-w -s"

ALL: lint build

.PHONY: lint
lint:
	./scripts/lint.sh

.PHONY: build
build:
	./scripts/build.sh
