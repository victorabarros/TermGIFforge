APP_NAME=$(shell pwd | xargs basename)
PWD=$(shell pwd)
API_APP_DIR=/go/src/github.com/victorabarros/${APP_NAME}
DOCKER_IMAGE=golang:1.23-alpine
BUILD_FILE=output/terminalGif

# build api
build-api:
	@echo "Building ${APP_NAME} to ./terminalGif"
	@docker run --rm -it \
		-v ${PWD}:${API_APP_DIR} -w ${API_APP_DIR} \
		--name ${APP_NAME} ${DOCKER_IMAGE} sh -c \
		"go build -o ${BUILD_FILE} cmd/server/main.go"
