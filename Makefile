CWD := $(shell pwd)
SRC_DIR := $(CWD)/envars
BIN_DIR := $(CWD)/bin
BIN_NAME = envars

FLAGS_BUILD := -o $(BIN_DIR)/$(BIN_NAME)

all: build.go

build.go: clear
	cd $(SRC_DIR); go build $(FLAGS_BUILD)
	clear

test.go:
	cd $(SRC_DIR); go test

clear:
	rm -rf $(BIN_DIR)/*
	clear