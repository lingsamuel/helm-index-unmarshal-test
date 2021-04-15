SHELL = /bin/bash
IMAGE_TAG := helm-index

build:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o output/main ./main.go
clean:
	rm -rf output
docker:
	DOCKER_BUILDKIT=1 docker build -f ./Dockerfile -t $(IMAGE_TAG) .
run: docker
	docker run --name helm-index --rm $(IMAGE_TAG)
