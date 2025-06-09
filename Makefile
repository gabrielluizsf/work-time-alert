PORT = 8089

build:
	@go tool wr -ico ./public/favicon.ico -manifest ./public/manifest.xml -o work-time-alert.syso
	@go build .
	@echo  "✅ Build completed"
test:
	@go test ./... -v
run: test build
	@PORT=$(PORT) go run ./runner/*.go
	@echo "http://localhost:$(PORT)"