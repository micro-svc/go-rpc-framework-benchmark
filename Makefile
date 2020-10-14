GOBUILD = go build -ldflags "-w -s"

ALL: lint build

.PHONY: lint
lint:
	./scripts/build.sh lint

.PHONY: build
build:
	./scripts/build.sh build

.PHONY: benchmark
benchmark:
	./scripts/benchmark.sh
