## Generate openapi.yaml
openapi-gen:
	@echo generating openapi.yaml file base on all specs files
	@swagger-cli bundle ./resources/open-api/openapi.yaml --outfile ./_build-openapi/openapi.yaml --type yaml

go-test:
	@echo Running tests
	@go test -v ./test/api_test.go

docker-build:
	@echo build docker
	@docker build --tag chat-service .

docker-down:
	@echo down docker
	@docker-compose down
	@docker rmi chat-service

docker-up:
	@echo up docker
	@docker build --tag chat-service .
	@docker-compose up

docker-exec:
	@docker exec -it mysql_container /bin/bash



