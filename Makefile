NAME=rafapi
.DEFAULT_GOAL:=list

go-test: # run tests
	./scripts/go-test.sh
.PHONY: test

go-lint: # run linter
	./scripts/go-lint.sh
.PHONY: lint

go-build: # build app
	./scripts/go-build.sh
.PHONY: compile

#docker-run: # run container
#	./scripts/docker-run.sh
#.PHONY: build

list: # list available commands
	@grep : ./Makefile | grep -v "grep\|.PHONY\|.DEFAULT_GOAL" | sed s/#/\\t#/
.PHONY: list

all: # test compile build and run
	@$(MAKE) go-lint
	@$(MAKE) go-test
	@$(MAKE) go-build
#	@$(MAKE) docker-run
.PHONY: all
