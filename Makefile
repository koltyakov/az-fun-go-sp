appname := az-fun-go-sp

start: build-local
	func start

debug:
	# Make fake server and shares the dynamic port which Functions are uses for custom handler session
	# Then Go server debug can be started in a usual way
	mkdir -p bin tmp && \
		echo "#!/bin/bash" > ./bin/server && \
		echo "echo \"Custom handler port: \$$FUNCTIONS_CUSTOMHANDLER_PORT\"" >> ./bin/server && \
		echo "printenv > ./tmp/.env" >> ./bin/server && \
		chmod +x ./bin/server
	func start

build-local:
	go build -o bin/server ./...

build-linux: clean
	GOOS=linux GOARCH=amd64 go build -v -ldflags "-s -w" -o bin/server ./...

publish: build-linux
	func azure functionapp publish $(appname)

clean:
	rm -rf bin/ tmp/