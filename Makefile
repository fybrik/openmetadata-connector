USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)

GIT_USER_ID := fybrik
GIT_REPO_ID := datacatalog-go
GIT_REPO_ID_MODELS := datacatalog-go-models
GIT_REPO_ID_CLIENT := datacatalog-go-client

FYBRIK_VERSION ?= v1.0.1

TMP_FILE = tmpfile.tmp

all: compile

.PHONY: compile
compile: patch
	go build .

.PHONY: run
run:
	go run . run --config conf/conf.yaml

.PHONY: generate-code
generate-code:
	if [ ! -d auto-generated/api ] || [ ! -d auto-generated/models ]; then \
	  git clone https://github.com/fybrik/fybrik/; \
	  cd fybrik && git checkout ${FYBRIK_VERSION}; \
	  docker run --rm \
             -v ${PWD}:/local \
             -u "${USER_ID}:${GROUP_ID}" \
             openapitools/openapi-generator-cli generate -g go-server \
             --additional-properties=serverPort=8081 \
             --git-user-id=${GIT_USER_ID} \
             --git-repo-id=${GIT_REPO_ID} \
             -o /local/auto-generated/api \
             -i /local/fybrik/connectors/api/datacatalog.spec.yaml; \
	  docker run --rm \
             -v ${PWD}:/local \
             -u "${USER_ID}:${GROUP_ID}" \
             openapitools/openapi-generator-cli generate -g go \
             --global-property=models,supportingFiles \
             --git-user-id=${GIT_USER_ID} \
             --git-repo-id=${GIT_REPO_ID_MODELS} \
             -o /local/auto-generated/models \
             -i /local/fybrik/connectors/api/datacatalog.spec.yaml; \
	  rm -Rf fybrik; \
	fi
	if [ ! -d auto-generated/client ]; then \
	  docker run --rm \
             -v ${PWD}:/local \
             -u "${USER_ID}:${GROUP_ID}" \
             openapitools/openapi-generator-cli generate -g go \
             --git-user-id=${GIT_USER_ID} \
             --git-repo-id=${GIT_REPO_ID_CLIENT} \
             -o /local/auto-generated/client \
             -i /local/client-swagger/swagger.json; \
	fi

.PHONY: patch
patch: generate-code
	awk '($$1 != "\"github.com/gorilla/mux\"") {print}' auto-generated/api/go/api_default.go > ${TMP_FILE}
	mv ${TMP_FILE} auto-generated/api/go/api_default.go
