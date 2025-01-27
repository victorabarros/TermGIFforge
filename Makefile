APP_NAME=vhsapi
PWD=$(shell pwd)
API_APP_DIR=/go/src/github.com/victorabarros/${APP_NAME}
DOCKER_IMAGE=golang:1.23-alpine
BUILD_FILE=output/terminalGif
PORT?=9001
COMMAND?="bash"

build-image:
	@echo "Building ${APP_NAME} image"
	@docker build --rm -t ${APP_NAME} .

debug-container:
	@echo "Debug ${APP_NAME} container on the port ${PORT}"
	@docker run --rm -it -p ${PORT}:80 \
		-v .:${API_APP_DIR} -w ${API_APP_DIR} \
		${APP_NAME} bash -c "${COMMAND}"

compile:
	@echo "Compiling ${APP_NAME} to ./main"
	@make debug-container COMMAND='go build cmd/server/main.go'
	@docker run --rm -it -p ${PORT}:80 \
		-v .:${API_APP_DIR} -w ${API_APP_DIR} \
		${APP_NAME} bash -c "${COMMAND}"
