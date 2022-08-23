USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)

GIT_USER_ID := fybrik
GIT_REPO_ID := datacatalog-go
GIT_REPO_ID_MODELS := datacatalog-go-models
GIT_REPO_ID_CLIENT := datacatalog-go-client

DOCKER_HOSTNAME ?= ghcr.io
DOCKER_NAMESPACE ?= fybrik
DOCKER_TAG ?= 0.0.0
DOCKER_NAME ?= openmetadata-connector

FYBRIK_VERSION ?= v1.0.1

TMP_FILE = tmpfile.tmp

IMG := ${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/${DOCKER_NAME}:${DOCKER_TAG}
export HELM_EXPERIMENTAL_OCI=1


all: generate-code patch

.PHONY: compile
compile:
	go build .

.PHONY: build
build: compile
	docker build . -t ${IMG}; cd ..

.PHONY: docker-push
docker-push:
	docker push ${IMG}

.PHONY: push-to-kind
push-to-kind:
	kind load docker-image ${IMG}

generate-code:
	git clone https://github.com/fybrik/fybrik/
	cd fybrik && git checkout ${FYBRIK_VERSION}
	docker run --rm \
           -v ${PWD}:/local \
           -u "${USER_ID}:${GROUP_ID}" \
           openapitools/openapi-generator-cli generate -g go-server \
           --additional-properties=serverPort=8081 \
           --git-user-id=${GIT_USER_ID} \
           --git-repo-id=${GIT_REPO_ID} \
           -o /local/api \
           -i /local/fybrik/connectors/api/datacatalog.spec.yaml
	docker run --rm \
           -v ${PWD}:/local \
           -u "${USER_ID}:${GROUP_ID}" \
           openapitools/openapi-generator-cli generate -g go \
           --global-property=models,supportingFiles \
           --git-user-id=${GIT_USER_ID} \
           --git-repo-id=${GIT_REPO_ID_MODELS} \
           -o /local/models \
           -i /local/fybrik/connectors/api/datacatalog.spec.yaml
	rm -Rf fybrik
	docker run --rm \
           -v ${PWD}:/local \
           -u "${USER_ID}:${GROUP_ID}" \
           openapitools/openapi-generator-cli generate -g go \
           --git-user-id=${GIT_USER_ID} \
           --git-repo-id=${GIT_REPO_ID_CLIENT} \
           -o /local/client \
           -i /local/client-swagger/swagger.json

patch:
	awk '($$1 != "\"github.com/gorilla/mux\"") {print}' api/go/api_default.go > ${TMP_FILE}
	mv ${TMP_FILE} api/go/api_default.go
