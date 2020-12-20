
VERSION                 ?= $(shell git describe --tags --always --dirty)
RELEASE_VERSION         ?= $(shell git describe --abbrev=0)
DOCKER_REPO							?= dereulenspiegel
DOCKER_IMAGE            ?= $(DOCKER_REPO)/taplist

LDFLAGS                	?= -X main.Version=$(VERSION) -w -s
GO_ENV                  = GO111MODULE=on CGO_ENABLED=0

GO_BUILD                = $(GO_ENV) go build -ldflags "$(LDFLAGS)"
GO_TEST                 = $(GO_ENV) go test -cover -v
PACKR_CMD								= packr2

.PHONY: clean build test gql-generate pack

pack:
	$(PACKR_CMD)

gql-generate:
	gqlgen generate

taplist: pack gql-generate
	$(GO_BUILD)

build: taplist

test:
	$(GO_TEST) ./...

clean: 
	rm -f ./taplist

rebuild: clean build
