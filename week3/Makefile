.PHONY: clean build release

clean:
	echo "clean"
	rm -rf bin

build: clean
	echo "building..."
	mkdir -p bin/amd64
	GOOS=linux GOARCH=amd64 go build -o bin/amd64/httpserver

release: build
	echo "release..."
	docker build -t httpserver:v1 .