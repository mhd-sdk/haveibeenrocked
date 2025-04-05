FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go module files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o backend cmd/main.go 

# Use a minimal base image for the final image
FROM alpine:latest

WORKDIR /app

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/backend /app/backend

# Command to run the application when the container starts
ENTRYPOINT ["/app/backend"]
