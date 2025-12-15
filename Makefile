APP_NAME=termgifforge
PWD=$(shell pwd)
WORK_DIR=/go/src/github.com/victorabarros/${APP_NAME}
IMAGE_NAME=${APP_NAME}-im
CONTAINER_NAME=${APP_NAME}
PORT?=9001
COMMAND?="bash"
ENV_FILE?=.env.local
BASE_IMAGE_NAME=golang:1.24.0
AUTOMATED_TESTS_PATH=zarf/automated-tests

build-image: remove-image
	@echo "Building ${IMAGE_NAME} image"
	@docker build --rm -t ${IMAGE_NAME} .

debug-container: kill-container
	@echo "Running ${APP_NAME} container on the port ${PORT}"
	@docker run -it -p ${PORT}:80 \
		--env-file ${ENV_FILE} --name ${CONTAINER_NAME} \
		-v ${PWD}:${WORK_DIR} -w ${WORK_DIR} \
		${IMAGE_NAME} bash -c "${COMMAND}"

compile: kill-container
	@echo "Compiling ${APP_NAME} to ./main"
	@rm -f ./main
	@make debug-container COMMAND="go build cmd/server/main.go"

run-app: kill-container
	@echo "Running ${APP_NAME} on the port ${PORT}"
	@make debug-container COMMAND="./main"

kill-container:
	@echo "Killing container ${CONTAINER_NAME}"
	@docker rm -f ${CONTAINER_NAME}

remove-image:
	@echo "Removing image ${IMAGE_NAME}"
	@docker rmi -f ${IMAGE_NAME}

tree:
	@docker container run --rm -it -v ${PWD}:${PWD} iankoulski/tree '-d ${PWD}' > TREE.md

debug-go-container:
	@echo "Running ${APP_NAME} container to run go commands"
	@docker run -it --rm --env-file ${ENV_FILE} --name ${CONTAINER_NAME}-go \
		-v ${PWD}:${WORK_DIR} -w ${WORK_DIR} \
		${BASE_IMAGE_NAME} bash -c "${COMMAND}"

test:
	@echo "Initalizing tests"
	@make debug-go-container COMMAND="\
		rm -f ${AUTOMATED_TESTS_PATH}/c.out && \
		go test ./... -v -cover -race -coverprofile=${AUTOMATED_TESTS_PATH}/c.out"

test-coverage:
	@echo "Building ${AUTOMATED_TESTS_PATH}/coverage.html"
	@rm -f ${AUTOMATED_TESTS_PATH}/coverage.html
	@make debug-go-container COMMAND="\
		go tool cover -html=${AUTOMATED_TESTS_PATH}/c.out -o ${AUTOMATED_TESTS_PATH}/coverage.html"

test-log:
	@echo "Writing ${AUTOMATED_TESTS_PATH}/tests.log"
	@rm -rf ${AUTOMATED_TESTS_PATH}/tests*.log
	@make test > ${AUTOMATED_TESTS_PATH}/tests.log
	@echo "Writing ${AUTOMATED_TESTS_PATH}/tests-summ.log"
	@cat ${AUTOMATED_TESTS_PATH}/tests.log  | grep "coverage: " > ${AUTOMATED_TESTS_PATH}/tests-summ.log
