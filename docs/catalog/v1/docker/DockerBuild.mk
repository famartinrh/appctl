
CONTAINER_IMAGE_TAG?=latest
CONTAINER_IMAGE_NAME?=quay.io/$(APP_ORGANIZATION)/$(APP_NAME):$(CONTAINER_IMAGE_TAG)

push:
	docker build -t $(CONTAINER_IMAGE_NAME) .