# Dockerfile

# Build stage
FROM golang:1.24.4-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .
RUN ls -l ./scripts

# Add execute permissions to the script
RUN chmod +x ./scripts/post_swag.sh

# Generate Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN /go/bin/swag init -g cmd/app/main.go && ls -l ./docs
RUN ./scripts/post_swag.sh

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main ./cmd/app

# Final stage
FROM alpine:latest
RUN apk add --no-cache curl

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy the generated Swagger docs
COPY --from=builder /app/docs ./docs

# Expose port 8080 to the outside world
EXPOSE 3080

# Command to run the executable
CMD ["./main"]