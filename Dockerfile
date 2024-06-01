FROM golang:1.22.3-alpine

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

CMD ["go", "run", "/app/server.go"]
