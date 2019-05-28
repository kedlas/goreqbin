
DOCKER_SERVICE_NAME := goreqbin
DOCKER_DEFAULT_TAG := latest
DOCKER_REGISTRY := kedlas/

format:
	gofmt -w cmd pkg

lint: format
	golint ./cmd/... ./pkg/...

test: lint
	go test ./...

run:
	go run cmd/reqbin.go

docker-build:
	docker build -t $(DOCKER_SERVICE_NAME):$(DOCKER_DEFAULT_TAG) -t $(DOCKER_REGISTRY)$(DOCKER_SERVICE_NAME):$(DOCKER_DEFAULT_TAG) .

docker-push:
	docker push $(DOCKER_REGISTRY)$(DOCKER_SERVICE_NAME):$(DOCKER_DEFAULT_TAG)

docker-run:
	docker run -it -p 8080:8080 -p 9090:9090 $(DOCKER_REGISTRY)$(DOCKER_SERVICE_NAME):$(DOCKER_DEFAULT_TAG)