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
	docker image build -t swagger-cli docker/swagger-cli && \
	docker run --rm -v $(PWD)/openapi:/app swagger-cli:latest bundle -o /app/combined.yaml /app/openapi.yaml && \
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
		-i /local/openapi/combined.yaml \
		-g go \
		-o /local/src/gen/openapi 