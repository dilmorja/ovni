TESTFLAGS	= -v -cover

test:
	@go test $(TESTFLAGS) ./...

test18:
	@go1.18 test $(TESTFLAGS) ./...