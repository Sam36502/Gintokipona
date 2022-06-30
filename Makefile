EXE_NAME := ginpona
BUILD_DIR := ./bin

build-win:
	@echo ' --> Building for Windows...'
	@GOOS=windows go build -o $(BUILD_DIR)/win/$(EXE_NAME).exe main.go

build-lin:
	@echo ' --> Building For Linux...'
	@GOOS=linux go build -o $(BUILD_DIR)/lin/$(EXE_NAME) main.go

build: build-win build-lin
	@echo ' --> Built all!'