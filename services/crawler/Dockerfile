FROM golang:1.18.1-alpine as development

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .

# Install Air Package for live reload
RUN go install github.com/cosmtrek/air@latest

# Start app
CMD air