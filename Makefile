NAME=rafapi
.DEFAULT_GOAL:=list

go-test: # run tests
	./scripts/go-test.sh
.PHONY: test

go-lint: # run linter
	./scripts/go-lint.sh
.PHONY: lint

go-build: # run compilation
	./scripts/go-build.sh
.PHONY: compile

docker-build: # build the app containers
	./scripts/docker-build.sh
.PHONY: docker-build

docker-start: # start the app containers
	./scripts/docker-start.sh
.PHONY: docker-start

docker-stop: # stop the app containers
	./scripts/docker-stop.sh
.PHONY: docker-stop

curl: # test the running application via curl
	./scripts/curl.sh
.PHONY: curl

list: # list available commands
	@grep : ./Makefile | grep -v "grep\|.PHONY\|.DEFAULT_GOAL" | sed s/#/\\t#/
.PHONY: list

all: # run the build pipeline and start the app
	@$(MAKE) go-build
	@$(MAKE) go-test
	@$(MAKE) go-lint
	@$(MAKE) docker-build
	@$(MAKE) docker-start
.PHONY: all
