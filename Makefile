start: build
	func start

build:
	go build -o bin/server ./...

build-linux: clean
	GOOS=linux GOARCH=amd64 go build -o bin/server ./...

clean:
	rm -rf bin/