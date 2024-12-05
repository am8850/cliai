default:
	@echo "Please specify a target to build"

build:
	@echo "Building the project..."
	go build 

publish: build
	@echo "Publishing the project..."	
	
	sudo cp cliaitemplates.json /usr/local/bin/cliaitemplates.json
	sudo cp cliaiopenai.json /usr/local/bin/cliaiopenai.json
	sudo cp cliai /usr/local/bin/cliai

	rm -f cliai