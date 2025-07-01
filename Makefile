build:
	@make swag
	@go build  -o bin/gahmen-api ./cmd/app/main.go

run: 
	@go run ./cmd/app/main.go

build-docker:
	CGO_ENABLED=0 GOOS=linux go build -o /gahmen-api ./cmd/app/main.go

test:
	@go test -v ./...

tidy:
	@go mod tidy

swag:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@~/go/bin/swag init -g cmd/app/main.go
	@./scripts/post_swag.sh