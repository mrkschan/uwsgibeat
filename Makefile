PREFIX?=/build

GOFILES = $(shell find . -type f -name '*.go')
uwsgibeat: $(GOFILES)
	go build

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm uwsgibeat || true

