
VERSION                 ?= $(shell git describe --tags --always --dirty)
RELEASE_VERSION         ?= $(shell git describe --abbrev=0)
DOCKER_REPO							?= dereulenspiegel
DOCKER_IMAGE            ?= $(DOCKER_REPO)/taplist

LDFLAGS                	?= -X main.Version=$(VERSION) -w -s
GO_ENV                  = GO111MODULE=on CGO_ENABLED=0

GO_BUILD                = $(GO_ENV) go build -ldflags "$(LDFLAGS)"
GO_TEST                 = $(GO_ENV) go test -cover -v
PACKR_CMD								= packr2

VUE_APP_GRAPHQL_HTTP		?= /query
# VUE_APP_GRAPHQL_WS			?= /query

.PHONY: clean build test gql-generate pack

frontend/index.html:
	cd frontend-src && VUE_APP_GRAPHQL_HTTP=$(VUE_APP_GRAPHQL_HTTP) VUE_APP_GRAPHQL_WS=$(VUE_APP_GRAPHQL_WS) yarn build

pack:
	$(PACKR_CMD)

gql-generate:
	gqlgen generate

taplist: frontend/index.html pack gql-generate
	$(GO_BUILD)

build: taplist

test:
	$(GO_TEST) ./...

clean: 
	rm -f ./taplist
	rm -rf ./frontend/*

rebuild: clean build
