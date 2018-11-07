APP_NAME = hello-handler
ZIP_NAME = $(APP_NAME).zip

HANDLER_DIR = ./cmd/$(APP_NAME)/
BUILD_DIR = build/
PROJECT_DIR = $(shell pwd)

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
dep:
	@echo dep
	@go get github.com/pelletier/go-toml
	@go get github.com/aws/aws-lambda-go/lambda
	@go get github.com/mike-neck/golang-lambda-deploy-example
resource:
	@echo resource
	@mkdir $(BUILD_DIR)
	@cp config.toml $(BUILD_DIR)
build: clean dep fmt vet resource
	@echo build
	@GOOS=linux go build -o $(BUILD_DIR)$(APP_NAME) $(HANDLER_DIR)main.go
archive: build
	@echo archive
	@cd $(BUILD_DIR) && \
	zip $(ZIP_NAME) $(APP_NAME) && \
	zip $(ZIP_NAME) config.toml && \
	cd $(PROJECT_DIR)
