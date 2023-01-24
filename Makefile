BINARY_NAME=packages-filter

all: build

build:
	go build -o $(BINARY_NAME)

run:
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)