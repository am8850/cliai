.PHONY: default build-linux build-windows build-macos publish

default:
	@echo "Please specify a target to build"

build:
	@echo "Building the project..."
	go build 

build-windows:
	@echo "Building the project for Windows..."
	GOOS=windows GOARCH=amd64 go build -o cliai.exe

build-macos:
	@echo "Building the project for MacOS..."
	GOOS=darwin GOARCH=amd64 go build -o cliai


	@echo "Installing team"
	sudo cp cliai /usr/local/bin/cliai
	sudo cp openai.json /usr/local/bin/cliaiopenai.json

dist: build
	rm -rf dist
	mkdir dist
	cp cliai dist/cliai
	cp cliaitemplates.json dist/cliaitemplates.json
	cp .env dist/cliaiopenai.json
	cd dist && zip cliai.zip cliai cliaitemplates.json cliaiopenai.json


install: build
	@echo "Publishing the project..."	
	
	sudo cp cliaitemplates.json /usr/local/bin/cliaitemplates.json
	sudo cp .env /usr/local/bin/cliaiopenai.json
	sudo cp cliai /usr/local/bin/cliai

	rm -f cliai