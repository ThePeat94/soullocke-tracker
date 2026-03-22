.PHONY: generate-sql export-openapi build run clean

generate-sql:
	sqlc generate

export-openapi:
	go run . export-openapi

build:
	mkdir -p dist
	go build -o dist/soullocke-backend .

run: build
	./dist/soullocke-backend

clean:
	rm -rf dist