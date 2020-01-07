PROJECT_NAME=golang_tutorial
VERSION=1.0.0

build:
	@go build -v -o $(PROJECT_NAME)-$(VERSION).bin .

install:
	@go install

test:
	@go test

docker_build_mysql:
	@docker-compose -f deployment/docker/docker-compose.yml up -d