APP_NAME = hello-handler
HANDLER_DIR = ./cmd/$(APP_NAME)/
BUILD_DIR = build/

fmt:
	@echo fmt
	@go fmt ./ ./cmd/...
vet:
	@echo vet
	@go vet ./ ./cmd/...
clean:
	@echo clean
	@go clean
	@rm -rf $(BUILD_DIR)
build: clean fmt vet
	@echo build
	@GOOS=linux go build -o $(BUILD_DIR)$(APP_NAME) $(HANDLER_DIR)main.go
