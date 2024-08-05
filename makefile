CMD=go
BIN_PATH=bin
SRC_PATH=.

all : build run

build:
	$(CMD) build -o $(BIN_PATH)/docker $(SRC_PATH)/main.go

clean:
	rm  $(BIN_PATH)/docker

run:
	sudo $(BIN_PATH)/docker run bash
