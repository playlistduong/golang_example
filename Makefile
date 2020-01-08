PROJECT_NAME=golang_example
VERSION=1.0.0
ENVIRONMENT_PREFIX=MYAPP

build:
	@go build -v -o $(PROJECT_NAME)-$(VERSION).bin .

install:
	@go install

test:
	@go test

docker_build_mysql:
	@docker-compose -f deployment/docker/docker-compose.yml up -d

environment:
	@export $(ENVIRONMENT_PREFIX)_DRIVER=mysql
	@export $(ENVIRONMENT_PREFIX)_USER=root
	@export $(ENVIRONMENT_PREFIX)_PASS=123456
	@export $(ENVIRONMENT_PREFIX)_DATABASE=test