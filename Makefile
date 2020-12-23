appname  := az-fun-go-sp
funcRoot := ./functions
srvPath  := ./functions/bin/server

start:
	@go build -tags "prod" -o $(funcRoot)/bin/server ./
	@cd $(funcRoot) && func start # --verbose

# Make fake server and shares the dynamic port which Functions use for custom handler session
# Then Go server debug can be started in a usual way
debug:
	@mkdir -p $(funcRoot)/bin $(funcRoot)/tmp && \
		echo "#!/bin/bash" > $(srvPath) && \
		echo "echo \"You should start custom handlers on http://127.0.0.1:\$$FUNCTIONS_CUSTOMHANDLER_PORT to debug.\"" >> $(srvPath) && \
		echo "echo \"This can be done by \\\"go run ./\\\" or VSCode Launch action.\"" >> $(srvPath) && \
		echo "printenv > ./tmp/.env" >> $(srvPath) && \
		chmod +x $(srvPath)
	@cd $(funcRoot) && func start

build:
	go build -o ./bin/server ./

build-prod: clean
	GOOS=linux GOARCH=amd64 go build -v -ldflags "-s -w" -tags "prod" -o $(funcRoot)/bin/server ./

publish: build-prod
	cd $(funcRoot) && func azure functionapp publish $(appname)

clean:
	rm -rf bin/ tmp/ $(funcRoot)/bin/ $(funcRoot)/tmp/
