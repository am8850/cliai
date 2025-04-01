.PHONY: default build-linux build-windows build-macos publish

default:
	@echo "Please specify a target to build"

build-linux:
	@echo "Building the project..."
	go build 

<<<<<<< HEAD
build-windows:
	@echo "Building the project for Windows..."
	GOOS=windows GOARCH=amd64 go build -o cliai.exe

build-macos:
	@echo "Building the project for MacOS..."
	GOOS=darwin GOARCH=amd64 go build -o cliai

publish: build-linux
=======
install: build
>>>>>>> 1c736dc (Add Docker command support and update Makefile target)
	@echo "Publishing the project..."	
	
	sudo cp cliaitemplates.json /usr/local/bin/cliaitemplates.json
	sudo cp cliaiopenai.json /usr/local/bin/cliaiopenai.json
	sudo cp cliai /usr/local/bin/cliai

	rm -f cliai