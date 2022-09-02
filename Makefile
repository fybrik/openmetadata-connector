USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)

GIT_HOST := fybrik.io
GIT_USER_ID := openmetadata-connector
GIT_REPO_ID := datacatalog-go
GIT_REPO_ID_MODELS := datacatalog-go-models
GIT_REPO_ID_CLIENT := datacatalog-go-client

FYBRIK_VERSION ?= v1.0.1

DOCKER_HOSTNAME ?= ghcr.io
DOCKER_NAMESPACE ?= fybrik
DOCKER_TAG ?= 0.0.0
DOCKER_NAME ?= openmetadata-connector

IMG := ${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/${DOCKER_NAME}:${DOCKER_TAG}
export HELM_EXPERIMENTAL_OCI=1

TMP_FILE = tmpfile.tmp

all: compile

.PHONY: compile
compile:
	go build .

.PHONY: run
run:
	go run . run --config conf/conf.yaml

# if datacatalog.spec.yaml or swagger.json changes in the fybrik repository,
# please remove the auto-generated dir and regenerate the code.
.PHONY: generate-code
generate-code:
	git clone https://github.com/fybrik/fybrik/
	cd fybrik && git checkout ${FYBRIK_VERSION}
	docker run --rm \
           -v ${PWD}:/local \
           -u "${USER_ID}:${GROUP_ID}" \
           openapitools/openapi-generator-cli generate -g go-server \
           --additional-properties=serverPort=8081 \
           --git-host=${GIT_HOST} \
           --git-user-id=${GIT_USER_ID} \
           --git-repo-id=${GIT_REPO_ID} \
           -o /local/auto-generated/api \
           -i /local/fybrik/connectors/api/datacatalog.spec.yaml
	docker run --rm \
           -v ${PWD}:/local \
           -u "${USER_ID}:${GROUP_ID}" \
           openapitools/openapi-generator-cli generate -g go \
           --global-property=models,supportingFiles \
           --git-host=${GIT_HOST} \
           --git-user-id=${GIT_USER_ID} \
           --git-repo-id=${GIT_REPO_ID_MODELS} \
           -o /local/auto-generated/models \
           -i /local/fybrik/connectors/api/datacatalog.spec.yaml
	rm -Rf fybrik
	docker run --rm \
           -v ${PWD}:/local \
           -u "${USER_ID}:${GROUP_ID}" \
           openapitools/openapi-generator-cli generate -g go \
           --git-host=${GIT_HOST} \
           --git-user-id=${GIT_USER_ID} \
           --git-repo-id=${GIT_REPO_ID_CLIENT} \
           -o /local/auto-generated/client \
           -i /local/client-swagger/swagger.json

.PHONY: patch
patch: generate-code
	awk '($$1 != "\"github.com/gorilla/mux\"") {print}' auto-generated/api/go/api_default.go > ${TMP_FILE}
	mv ${TMP_FILE} auto-generated/api/go/api_default.go

.PHONY: build
build: compile
	docker build . -t ${IMG}; cd ..

.PHONY: docker-push
docker-push:
	docker push ${IMG}

.PHONY: push-to-kind
push-to-kind:
	kind load docker-image ${IMG}
