NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

# The import path is the unique absolute name of your repository.
# All subpackages should always be imported as relative to it.
# If you change this, run `make clean`.
PKG_SRC := github.com/italolelis/reachable

.PHONY: all clean setup build

all: clean setup test build

setup:
	@echo "$(OK_COLOR)==> Installing dependencies$(NO_COLOR)"
	@go get -u github.com/golang/dep/cmd/dep
	@go get -u github.com/golang/lint/golint
	@dep ensure -vendor-only

build:
	@echo "$(OK_COLOR)==> Building... $(NO_COLOR)"
	CGO_ENABLED=0 go build -ldflags "-s -w" -ldflags "-X cmd.version=$(VERSION)" -o "dist/reachable" $(PKG_SRC)

test: lint format vet
	@echo "$(OK_COLOR)==> Running tests$(NO_COLOR)"
	@go test -v -cover ./...

format:
	@echo "$(OK_COLOR)==> checking code formating with 'gofmt' tool$(NO_COLOR)"
	@gofmt -l -s cmd pkg | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

vet:
	@echo "$(OK_COLOR)==> checking code correctness with 'go vet' tool$(NO_COLOR)"
	@go vet ./...

lint: tools.golint
	@echo "$(OK_COLOR)==> checking code style with 'golint' tool$(NO_COLOR)"
	@go list ./... | xargs -n 1 golint -set_exit_status

clean:
	@echo "$(OK_COLOR)==> Cleaning project$(NO_COLOR)"
	@go clean
	@rm -rf bin $GOPATH/bin

#---------------
#-- tools
#---------------

.PHONY: tools tools.dep tools.golint
tools: tools.dep tools.golint

tools.golint:
	@command -v golint >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing golint"; \
		go get github.com/golang/lint/golint; \
	fi

tools.dep:
	@command -v dep >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing dep"; \
		@go get -u github.com/golang/dep/cmd/dep; \
	fi