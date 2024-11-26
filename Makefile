default:
	@echo "Please specify a target to build"

build:
	@echo "Building the project..."
	go build 

publish: build
	@echo "Publishing the project..."	
	sudo cp cliai /usr/local/bin/cliai
	sudo cp openai.json /usr/local/bin/openai.json
	rm -f cliai