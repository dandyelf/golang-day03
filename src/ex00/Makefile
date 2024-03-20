BUILD_DIR := build
TARGET := myElasticCli

# curl -s -XGET "http://localhost:9200/places/_doc/1"

all: clean build test

build:
	go build -o $(BUILD_DIR)/$(TARGET) main.go csv_db.go

test:
	@echo '\n'\*************$(TARGET) test***************************'\n'
	@echo Find in Elastic
	./$(BUILD_DIR)/$(TARGET)
	@echo
	@echo \****************************************

run:
	./$(BUILD_DIR)/$(TARGET) ./tests

clean:
	rm -rf $(BUILD_DIR)
