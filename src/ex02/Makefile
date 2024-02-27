BUILD_DIR := build
TARGET := myXargs

all: clean build test

build:
	go build -o $(BUILD_DIR)/$(TARGET) main.go

test: myFind myWc
	mv ../ex00/build/myFind ./build/myFind
	mv ../ex01/build/myWc ./build/myWc
	@echo '\n'\*************myXargs test***************************'\n'

	cd build &&	./myFind -f -ext 'txt' ../tests | ./myXargs ./myWc -l

	@echo '\n'\****************************************************'\n'

run:
	./$(BUILD_DIR)/$(TARGET)

clean:
	rm -rf $(BUILD_DIR)

myFind:
	cd ../ex00 && make

myWc:
	cd ../ex01 && make

