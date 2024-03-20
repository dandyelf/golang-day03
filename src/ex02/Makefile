BUILD_DIR := build
TARGET := myXargs

all: clean build test

build:
	go build -o $(BUILD_DIR)/$(TARGET) main.go

test:
	@echo '\n'\*************myXargs test***************************'\n'


	cd build &&	./myFind -f -ext 'txt' ../tests | ./myXargs ./myWc -l

	@echo '\n'\****************************************************'\n'

run:
	./$(BUILD_DIR)/$(TARGET)

clean:
	rm -rf $(BUILD_DIR)

myFind:
	cd ../ex00 && make
