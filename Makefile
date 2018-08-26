NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

.PHONY: all clean build

all: clean test build

build:
	@echo "$(OK_COLOR)==> Building... $(NO_COLOR)"
	go build -ldflags "-s -w" -ldflags "-X cmd.version=$(VERSION)" -o "dist/reachable" .

test: format vet
	@echo "$(OK_COLOR)==> Running tests$(NO_COLOR)"
	@go test -v -cover ./...

format:
	@echo "$(OK_COLOR)==> checking code formating with 'gofmt' tool$(NO_COLOR)"
	@gofmt -l -s cmd pkg | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

vet:
	@echo "$(OK_COLOR)==> checking code correctness with 'go vet' tool$(NO_COLOR)"
	@go vet ./...

clean:
	@echo "$(OK_COLOR)==> Cleaning project$(NO_COLOR)"
	@go clean
