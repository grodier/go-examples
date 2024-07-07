.PHONY:example-1
example-1:
	@echo "Simple GET"
	go run examples/simple-get/main.go

.PHONEY:example-2
example-2:
	@echo "GET w/ NewRequestWithContext and Do"
	go run examples/get-with-newrequestwithcontext/main.go
