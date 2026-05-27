.PHONY: *

build:
	go build

fmt:
	go fmt .

fmt-check:
	@test -z "$$(gofmt -l .)" || (gofmt -l . && exit 1)

