.PHONY: test
test:
	go test -race -v ./...

.PHONY: style-fix
style-fix:
	gofmt -w .

.PHONE: lint
lint:
	golangci-lint run
