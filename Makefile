APP_NAME=termgifforge
PWD=$(shell pwd)
WORK_DIR=/go/src/github.com/victorabarros/${APP_NAME}
IMAGE_NAME=${APP_NAME}-im
CONTAINER_NAME=${APP_NAME}
PORT?=9001
COMMAND?="bash"

build-image:
	@echo "Building ${IMAGE_NAME} image"
	@docker build --rm -t ${IMAGE_NAME} .

debug-container:
	@echo "Debug ${APP_NAME} container on the port ${PORT}"
	@docker run --rm -it -p ${PORT}:80 \
		--env ENVIRONMENT=local --name ${CONTAINER_NAME} \
		-v ${PWD}:${WORK_DIR} -w ${WORK_DIR} \
		${IMAGE_NAME} bash -c "${COMMAND}"

compile:
	@echo "Compiling ${APP_NAME} to ./main"
	@rm -f ./main
	@docker run --rm \
		--env ENVIRONMENT=local --name ${CONTAINER_NAME} \
		-v ${PWD}:${WORK_DIR} -w ${WORK_DIR} \
		${IMAGE_NAME} bash -c "go build cmd/server/main.go"

run-app: kill-container
	@echo "Running ${APP_NAME} on the port ${PORT}"
	@docker run --rm -d -p ${PORT}:80 \
		--env ENVIRONMENT=local --name ${CONTAINER_NAME} \
		-v ${PWD}:${WORK_DIR} -w ${WORK_DIR} \
		${IMAGE_NAME} bash -c "./main"

kill-container:
	@echo "Killing container ${CONTAINER_NAME}"
	@docker rm -f ${CONTAINER_NAME}
