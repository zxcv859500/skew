ifndef BUILD_SCOPE
BUILD_SCOPE=dev
endif

.PHONY: all test clean build run build-dev build-local

CMD_PROJECT=./cli

build:
	@echo "Build option: local"
	@go build -o ./bin/entry-point $(CMD_PROJECT)

run: build
	@./bin/entry-point apply -f ./testdata/ParseThisYaml.yaml --warnings-as-errors