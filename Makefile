include .env

all: test vet fmt lint build

dev:
	find . | grep '\.go' | entr -r go run .

clients:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/openapi/evermile.yaml \
    -g go \
    --additional-properties=isGoSubmodule=true \
    --additional-properties=packageName=EvermileClient \
    -o /local/evermile-client

ngrok:
	docker run -it -e NGROK_AUTHTOKEN=${NGROK_AUTHTOKEN} ngrok/ngrok:latest http host.docker.internal:${HTTP_PORT}

test:
	go test ./...

vet:
	go vet ./...

fmt:
	go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l
	test -z $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l)

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

build:
	go build -o deliveries ./...

emToken:
	@echo -n ${EVERMILE_CLIENT_ID}:${EVERMILE_CLIENT_SECRET} | base64
