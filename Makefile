GOBUILD = go build -ldflags "-w -s"

ALL: lint build

.PHONY: gomod-tidy
gomod-tidy:
	./scripts/build.sh gomod-tidy

.PHONY: lint
lint:
	./scripts/build.sh lint

.PHONY: build
build:
	./scripts/build.sh build

.PHONY: benchmark
benchmark:
	./scripts/benchmark.sh

.PHONY: gen-docs
gen-docs:
	cd tools && go run gen/main.go --gen=docs >../docs/index.html 2>&1

.PHONY: gen-readme
gen-readme:
	cd tools && go run gen/main.go --gen=readme >../README.md 2>&1
