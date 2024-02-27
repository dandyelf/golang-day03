BUILD_DIR := build
TARGET := myRotation
FILESDIR := ../tests

all: clean build test

build:
	go build -o $(BUILD_DIR)/$(TARGET) main.go

test:
	@echo '\n'\*************$(TARGET) test***************************'\n'

	cd $(BUILD_DIR) && ./$(TARGET) $(FILESDIR)/some_application.txt
	@echo '\n'*
	cd $(BUILD_DIR) && ./$(TARGET) $(FILESDIR)/some_application.txt $(FILESDIR)/other_application.txt
	@echo '\n'*
	cd $(BUILD_DIR) && ./$(TARGET) -a /data/archive $(FILESDIR)/some_application.txt
	@echo '\n'*
	cd $(BUILD_DIR) && ./$(TARGET) -a /data/archive $(FILESDIR)/some_application.txt $(FILESDIR)/other_application.txt
	@echo '\n'*
	cd $(BUILD_DIR) && ./$(TARGET)
	@echo '\n'*
	cd $(BUILD_DIR) && ./$(TARGET) -a
	@echo '\n'*
	cd $(BUILD_DIR) && ./$(TARGET) -a /data/archive

	@echo '\n'\****************************************************'\n'

run:
	./$(BUILD_DIR)/$(TARGET)

clean:
	rm -rf $(BUILD_DIR)
