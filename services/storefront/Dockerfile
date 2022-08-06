FROM golang:1.18.1-alpine as development

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .

# Start app
CMD go run ./cmd/main.go
