.PHONY: default build-linux build-windows build-macos publish

default:
	@echo "Please specify a target to build"

build-linux:
	@echo "Building the project..."
	go build 

install: build
	@echo "Publishing the project..."	
	
	sudo cp cliaitemplates.json /usr/local/bin/cliaitemplates.json
	sudo cp cliaiopenai.json /usr/local/bin/cliaiopenai.json
	sudo cp cliai /usr/local/bin/cliai

	rm -f cliai