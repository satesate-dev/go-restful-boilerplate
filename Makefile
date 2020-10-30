GOCMD	= go
GOBUILD	= $(GOCMD) build
GOCLEAN	= $(GOCMD) clean
GOTEST	= $(GOCMD) test
GOGET	= $(GOCMD) get
BINARY_NAME	= server
LINUX_BUILD = GOOS=linux GOARCH=arm GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
TAG = server/v0.0.1 # Please change the name tag if you want to deploy
DOCKER_BUILD = docker build -t $(TAG) .
PWD = $(shell pwd)
DOCKER_RUN_LOCAL = docker run -d --restart=always --name=$(BINARY_NAME) -p 1111:1111 -v $(PWD)/.config.toml:/root/.config.toml --mount type=bind,source=/var/log/$(BINARY_NAME),destination=/root/log $(TAG)
CLEAN_UP = rm -rf $(BINARY_NAME) # Remove unused binary

# build command will automatic detect your os
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# build and run binary
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)

# build for linux os
build-linux:
	$(LINUX_BUILD)

# deploy based on alpine linux
deploy:
	$(CHMODGOSUM)
	$(LINUX_BUILD)
	$(DOCKER_BUILD)
	$(DOCKER_RUN_LOCAL)
	$(CLEAN_UP)
