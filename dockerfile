# Stage 1: Build the Go binary
FROM golang:1.23.1-bullseye AS builder
WORKDIR /app

# Copy go.mod files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o myapp main.go

# Stage 2: Create a minimal image with just the binary
FROM gcr.io/distroless/base-debian10
WORKDIR /app

# Copy the Go binary and templates from the builder stage
COPY --from=builder /app/myapp .
COPY --from=builder /app/templates /app/templates

# Expose the port the app listens on
EXPOSE 8080

# Set the entry point to the binary
ENTRYPOINT ["./myapp"]


