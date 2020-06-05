IMAGE_NAME=dynred
TAG=0.0.1
NOTE_USER_ID=your_id

build:
	docker build -t ${IMAGE_NAME}:${TAG} .
	docker build -t ${IMAGE_NAME}:latest .
.PHONY: build

run:
	docker run -it -p 8080:8080 --env NOTE_USER_ID=${NOTE_USER_ID} ${IMAGE_NAME}:${TAG}
