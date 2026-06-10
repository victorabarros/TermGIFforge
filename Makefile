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
	@docker run -d -p ${PORT}:80 \
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

debug-go-container:
	@docker run -it --rm --env-file ${ENV_FILE} --name ${CONTAINER_NAME}-go \
		-v ${PWD}:${WORK_DIR} -w ${WORK_DIR} \
		${BASE_IMAGE_NAME} bash -c "${COMMAND}"

debug-python-container:
	@docker run -it --rm --env-file ${ENV_FILE} --name ${CONTAINER_NAME}-python \
		-v ${PWD}:${WORK_DIR} -w ${WORK_DIR} \
		python:alpine sh

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

# Infrastructure

connect-vm:
	@ssh hostinger

# Git

commit-llm-generated:
	@msg_file="$$(mktemp)"; \
	{ \
		printf '%s\n\n' 'Write the final git commit message for the staged changes.'; \
		printf '%s\n' 'Return only the commit message text that should be passed to git commit.'; \
		printf '%s\n' 'Do not repeat these instructions.'; \
		printf '%s\n' 'Do not include markdown, code examples, code fences, labels, quotes, explanations, or diff summaries.'; \
		printf '%s\n' 'Use imperative mood.'; \
		printf '%s\n' 'Keep the subject line under 72 characters.'; \
		printf '%s\n' 'Add a short body only if it materially improves clarity.'; \
		printf '%s\n' 'If there is a body, separate it from the subject with one blank line.'; \
		printf '\n%s\n' 'git status --short:'; \
		git status --short; \
		printf '\n%s\n' 'git diff --cached --stat:'; \
		git diff --cached --stat; \
		printf '\n%s\n' 'git diff --cached:'; \
		git diff --cached; \
	} | ollama run "$(OLLAMA_MODEL)" > "$$msg_file"; \
	printf '🦙 ollama generated' >> "$$msg_file"; \
	printf '%s\n' 'Generated commit message:'; \
	cat "$$msg_file"; \
	printf '\n'; \
	commit_msg="$$(perl -pe 's/\e\[[0-9;?]*[ -\/]*[@-~]//g' "$$msg_file")"; \
	rm -f "$$msg_file"; \
	git commit -m "$$commit_msg"

push p:
	git add .
	make commit-llm-generated
	git push
