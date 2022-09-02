USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)

GIT_HOST := fybrik.io
GIT_USER_ID := openmetadata-connector
GIT_REPO_ID := datacatalog-go
GIT_REPO_ID_MODELS := datacatalog-go-models
GIT_REPO_ID_CLIENT := datacatalog-go-client

AUTO_GENERATED = auto-generated
FYBRIK_VERSION ?= v1.0.1

TMP_FILE = tmpfile.tmp

all: compile

.PHONY: compile
compile:
	go build .

.PHONY: run
run:
	go run . run --config conf/conf.yaml

# if datacatalog.spec.yaml or swagger.json changes in the fybrik repository,
.PHONY: generate-code
generate-code:
	rm -Rf fybrik
	rm -Rf ${AUTO_GENERATED}
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
           -o /local/${AUTO_GENERATED}/api \
           -i /local/fybrik/connectors/api/datacatalog.spec.yaml
	docker run --rm \
           -v ${PWD}:/local \
           -u "${USER_ID}:${GROUP_ID}" \
           openapitools/openapi-generator-cli generate -g go \
           --global-property=models,supportingFiles \
           --git-host=${GIT_HOST} \
           --git-user-id=${GIT_USER_ID} \
           --git-repo-id=${GIT_REPO_ID_MODELS} \
           -o /local/${AUTO_GENERATED}/models \
           -i /local/fybrik/connectors/api/datacatalog.spec.yaml
	rm -Rf fybrik
	docker run --rm \
           -v ${PWD}:/local \
           -u "${USER_ID}:${GROUP_ID}" \
           openapitools/openapi-generator-cli generate -g go \
           --git-host=${GIT_HOST} \
           --git-user-id=${GIT_USER_ID} \
           --git-repo-id=${GIT_REPO_ID_CLIENT} \
           -o /local/${AUTO_GENERATED}/client \
           -i /local/client-swagger/swagger.json

.PHONY: patch
patch: generate-code
	awk '($$1 != "\"github.com/gorilla/mux\"") {print}' ${AUTO_GENERATED}/api/go/api_default.go > ${TMP_FILE}
	mv ${TMP_FILE} ${AUTO_GENERATED}/api/go/api_default.go
