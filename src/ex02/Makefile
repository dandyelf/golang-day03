BUILD_DIR := build
TARGET := myJson
OS = $(shell uname)

all: build test

build: clean
	go build -o $(BUILD_DIR)/$(TARGET)

test:
	cd ../ex00 && make && make clean
	@echo '\n'\************* $(TARGET) ***************************'\n'
	@echo
	@echo Push elasticsearch data
ifeq ($(OS), Darwin)
	open http://127.0.0.1:8888/api/places?page=3
else
	@echo Server start on addres: http://127.0.0.1:8888/api/places?page=3
endif
	./$(BUILD_DIR)/$(TARGET)
	@echo
	@echo \****************************************

run: build
	@echo Server start on addres: http://127.0.0.1:8888/api/places?page=3
ifeq ($(OS), Darwin)
	open http://127.0.0.1:8888/api/places?page=3
endif
	./$(BUILD_DIR)/$(TARGET)

clean:
	rm -rf $(BUILD_DIR)
