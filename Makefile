
# load env vars
-include .env
export API := $(value API)

.PHONY: run
run:
	go run .

.PHONY: test
test:
	go test -v ./...