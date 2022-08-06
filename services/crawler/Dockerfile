FROM golang:1.18.1-alpine as development

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .

# Install Reflex for development
RUN go install github.com/cespare/reflex@latest

# Start app
CMD reflex -g './**/*.go' go run main.go --start-service
