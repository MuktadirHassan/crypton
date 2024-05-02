BINARY_NAME=crypton

build:
	go build -o bin/$(BINARY_NAME) -v

run: build
	./bin/$(BINARY_NAME)

build-docker:
	docker build -t $(BINARY_NAME) .

clean:
	go clean
	rm -f bin/$(BINARY_NAME)