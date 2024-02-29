BUILD_DIR := build
TARGET := myElasticCli

all: clean build test

build:
	go build -o $(BUILD_DIR)/$(TARGET) main.go

test:
	@echo '\n'\*************myFind test***************************'\n'
	@echo Find in Elastic
	./$(BUILD_DIR)/$(TARGET)
	@echo
	@echo \****************************************

run:
	./$(BUILD_DIR)/$(TARGET) ./tests

clean:
	rm -rf $(BUILD_DIR)
