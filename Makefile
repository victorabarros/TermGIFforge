APP_NAME=termgifforge
PWD=$(shell pwd)
WORK_DIR=/go/src/github.com/victorabarros/${APP_NAME}
IMAGE_NAME=${APP_NAME}-im
CONTAINER_NAME=${APP_NAME}
PORT?=9001
COMMAND?="bash"
ENV_FILE?=.env.local

build-image:
	@echo "Building ${IMAGE_NAME} image"
	@docker build --rm -t ${IMAGE_NAME} .

debug-container:
	@echo "Debug ${APP_NAME} container on the port ${PORT}"
	@docker run --rm -it -p ${PORT}:80 \
		--env-file ${ENV_FILE} --name ${CONTAINER_NAME} \
		-v ${PWD}:${WORK_DIR} -w ${WORK_DIR} \
		${IMAGE_NAME} bash -c "${COMMAND}"

compile: kill-container
	@echo "Compiling ${APP_NAME} to ./main"
	@rm -f ./main
	@docker run --rm --name ${CONTAINER_NAME} \
		-v ${PWD}:${WORK_DIR} -w ${WORK_DIR} \
		${IMAGE_NAME} bash -c "go build cmd/server/main.go"

run-app: kill-container
	@echo "Running ${APP_NAME} on the port ${PORT}"
	@docker run --rm -d -p ${PORT}:80 \
		--env-file ${ENV_FILE} --name ${CONTAINER_NAME} \
		-v ${PWD}:${WORK_DIR} -w ${WORK_DIR} \
		${IMAGE_NAME} bash -c "./main"

kill-container:
	@echo "Killing container ${CONTAINER_NAME}"
	@docker rm -f ${CONTAINER_NAME}

remove-image:
	@echo "Removing image ${IMAGE_NAME}"
	@docker rmi -f ${IMAGE_NAME}

tree:
	@docker container run --rm -it -v ${PWD}:${PWD} iankoulski/tree '-d ${PWD}' > TREE.md
