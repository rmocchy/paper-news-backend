.PHONY: install-tools
install-tools:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: build
build:
	cd src && \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../build/main

.PHONY: start
start:
	docker compose up -d --build

.PHONY: openapi-gen
openapi-gen:
	truncate -s 0 openapi/combined.yaml && \
	rm -rf src/gen/openapi && \
	docker image build -t swagger-cli config/swagger-cli && \
	docker run --rm -v $(PWD)/openapi:/app swagger-cli:latest bundle -o /app/combined.yaml /app/openapi.yaml && \
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
		-i /local/openapi/combined.yaml \
		-g go \
		-o /local/src/gen/openapi \
		--global-property=models,apis,supportingFiles=configuration.go:client.go:utils.go,apiTests=false,apiDocs=false,modelTests=false,modelDocs=false

.PHONY: oapi-codegen
oapi-codegen:
	truncate -s 0 openapi/combined.yaml && \
	rm -rf gen/openapi/* && \
	docker image build -t swagger-cli config/swagger-cli && \
	docker run --rm -v $(PWD)/openapi:/app swagger-cli:latest bundle -o /app/combined.yaml /app/openapi.yaml && \
	oapi-codegen -config config/oapi-codegen/config.yaml -o gen/openapi/openapi.gen.go openapi/combined.yaml