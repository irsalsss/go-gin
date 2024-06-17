# Stage 1: Build the Go application
FROM golang:1.22.3 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /build

# Stage 2: Create a small image with the compiled Go binary
FROM alpine:latest

# Install necessary CA certificates for HTTPS connections
RUN apk add --no-cache ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /build /build

# Expose the port on which the app will run
EXPOSE 8080

# Command to run the application
CMD ["/build"]