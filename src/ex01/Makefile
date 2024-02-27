BUILD_DIR := build
TARGET := myWc

all: clean build test

build:
	go build -o $(BUILD_DIR)/$(TARGET) main.go

test:
	@echo '\n'\*************myWc test***************************'\n'
	@echo Flag -l '(' line feed is on the previous line ')'
	./$(BUILD_DIR)/$(TARGET) -l ./Makefile
	@echo
	@echo Flag -m
	./$(BUILD_DIR)/$(TARGET) -m ./tests/aaa.txt
	@echo
	@echo Flag -w
	./$(BUILD_DIR)/$(TARGET) -w ./tests/aaa.txt
	@echo
	@echo NoFlag
	./$(BUILD_DIR)/$(TARGET) ./tests/aaa.txt ./tests/bbb.txt
	@echo
	@echo Flag cros -w -l
	./$(BUILD_DIR)/$(TARGET) -w -l Makefile
	@echo
	@echo \****************************************

run:
	./$(BUILD_DIR)/$(TARGET) ./tests/aaa.txt

clean:
	rm -rf $(BUILD_DIR)
