GOBUILD = go build -ldflags "-w -s"

ALL: lint build benchmark

.PHONY: lint
lint:
	./scripts/lint.sh

.PHONY: build
build:
	./scripts/build.sh

.PHONY: benchmark
benchmark:
	./scripts/benchmark.sh
