APP_NAME=$(shell pwd | xargs basename)
PWD=$(shell pwd)
API_APP_DIR=/go/src/github.com/victorabarros/${APP_NAME}
DOCKER_IMAGE=golang:1.23-alpine

# build api
build-api:
	@echo "Building ${APP_NAME} to ./terminalGif"
	@docker rm -f ${APP_NAME}
	@docker run -it \
		-v ${PWD}:${API_APP_DIR} -w ${API_APP_DIR} \
		--name ${APP_NAME} ${DOCKER_IMAGE} sh -c \
		"go build -o terminalGif cmd/server/main.go"
