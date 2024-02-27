BUILD_DIR := build
TARGET := myFind

all: clean build test

build:
	go build -o $(BUILD_DIR)/$(TARGET) main.go

test:
	@echo '\n'\*************myFind test***************************'\n'
	@echo Flag -d
	./$(BUILD_DIR)/$(TARGET) -d ./tests
	@echo
	@echo Flag -f
	./$(BUILD_DIR)/$(TARGET) -f ./tests
	@echo
	@echo Flag -sl
	./$(BUILD_DIR)/$(TARGET) -f -sl ./tests
	@echo
	@echo Flag -ext
	./$(BUILD_DIR)/$(TARGET) -f  -ext 'nc' ./tests
	@echo
	@echo Flag cros -f -d
	./$(BUILD_DIR)/$(TARGET) -f -d ./tests
	@echo
	@echo Flag cros -f -d -ext
	./$(BUILD_DIR)/$(TARGET) -f -d -ext 'nc' ./tests
	@echo
	@echo Flag -ext wt -f
	./$(BUILD_DIR)/$(TARGET) -ext 'txt' ./tests

	@echo \****************************************

run:
	./$(BUILD_DIR)/$(TARGET) ./tests

clean:
	rm -rf $(BUILD_DIR)
