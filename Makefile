go-rpc-service:
	env GO111MODULE=on go build $(LDFLAGS)
.PHONY: go-rpc-service

clean:
	rm go-rpc-service

test:
	go test -v ./...

lint:
	golangci-lint run ./...