IMAGE_NAME=tracking/dynred
TAG=0.0.9
NOTE_USER_ID=your_id
PROJECT_ID=soto-goto

build:
	docker build -t ${IMAGE_NAME}:${TAG} .
	docker build -t ${IMAGE_NAME}:latest .
.PHONY: build

push: build
	docker tag ${IMAGE_NAME}:${TAG} asia.gcr.io/${PROJECT_ID}/${IMAGE_NAME}:${TAG}
	docker push asia.gcr.io/${PROJECT_ID}/${IMAGE_NAME}:${TAG}
.PHONY: push

run:
	docker run -it -p 8080:8080 --env NOTE_USER_ID=${NOTE_USER_ID} ${IMAGE_NAME}:${TAG}
