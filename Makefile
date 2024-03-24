build:
	@go build  -o bin/gahmen-api ./cmd/app/main.go

run: 
	@go run ./cmd/app/main.go

test:
	@go test -v ./...

tidy:
	@go mod tidy	