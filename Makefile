.PHONY: default build-linux build-windows build-macos publish

default:
	@echo "Please specify a target to build"

build:
	@echo "Building the project..."
	go build 

dist-windows:
	@echo "Building the project for Windows..."
	GOOS=windows GOARCH=amd64 go build -o cliai.exe
	rm -rf dist && mkdir dist
	sudo cp cliaitemplates.json dist/cliaitemplates.json
	sudo cp .env dist/cliaiopenai.json
	sudo cp cliai.exe dist/cliai.exe

dist-macos:
	@echo "Building the project for MacOS..."
	GOOS=darwin GOARCH=amd64 go build -o cliai	
	rm -rf dist && mkdir dist
	sudo cp cliaitemplates.json dist/cliaitemplates.json
	sudo cp .env dist/cliaiopenai.json
	sudo cp cliai dist/cliai

dist: build
	rm -rf dist
	mkdir dist
	cp cliai dist/cliai
	cp cliaitemplates.json dist/cliaitemplates.json
	cp cliaiopenai.json dist/cliaiopenai.json
	cd dist && zip cliai.zip cliai cliaitemplates.json cliaiopenai.json
	cd dist && rm -f cliai cliaitemplates.json cliaiopenai.json	

install: build
	@echo "Publishing the project..."		
	sudo cp cliaitemplates.json /usr/local/bin/cliaitemplates.json
	sudo cp cliaiopenai.json /usr/local/bin/cliaiopenai.json
	sudo cp cliai /usr/local/bin/cliai
	rm -f cliai