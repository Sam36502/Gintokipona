EXE_NAME := ginpona
BUILD_DIR := ./bin

build-win:
	GOOS=windows go build -o $(BUILD_DIR)/win/$(EXE_NAME).exe main.go

build-lin:
	GOOS=linux go build -o $(BUILD_DIR)/lin/$(EXE_NAME) main.go