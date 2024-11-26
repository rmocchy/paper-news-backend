.PHONY: build
build:
	cd src && \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../build/main

start:
	docker compose up -d --build