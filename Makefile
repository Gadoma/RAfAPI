NAME=rafapi
.DEFAULT_GOAL := list

go-test: # run tests
	./scripts/test.sh
.PHONY: test

go-lint: # run linter
	./scripts/lint.sh
.PHONY: lint

go-compile: # compile app for Linux@AMD64 on OSX@ARM64 via MUSL
	./scripts/compile.sh
.PHONY: compile

docker-build: # build container
	./scripts/build.sh
.PHONY: build

docker-run: # run container
	./scripts/run.sh
.PHONY: build

clean: # delete generated files
	./scripts/clean.sh
.PHONY: compile

list: # list available commands
	@grep : ./Makefile | grep -v "grep\|.PHONY\|.DEFAULT_GOAL" | sed s/#/\\t#/
.PHONY: list

all: # test compile build and run
	@$(MAKE) go-test
	@$(MAKE) clean
	@$(MAKE) go-compile
	@$(MAKE) docker-build
	@$(MAKE) docker-run
.PHONY: all

