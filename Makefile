build:
	@go build .
	@echo  "Build completed"
test:
	@go test ./... -v