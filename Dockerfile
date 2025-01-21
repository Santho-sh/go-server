# Start from golang base image
FROM golang:1.22-alpine as builder

# Set working directory
WORKDIR /app

# Copy go mod file first
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .
COPY data.json .

# Expose port 4000
EXPOSE 4000

# Command to run the executable
CMD ["./main"]