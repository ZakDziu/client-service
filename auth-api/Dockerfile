FROM golang:1.22.6 as development

# Add a work directory
WORKDIR /app

# Install packages
RUN apt-get update && \
    apt-get install -y nano htop

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Set timezone
ENV TZ=UTC
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Copy app files
COPY . .
RUN go build -o ./app ./cmd/main.go

# Install Reflex for development
RUN go install github.com/cespare/reflex@latest

# Expose port
EXPOSE 8081

# Start app
CMD reflex -r '\.go$' go run cmd/main.go --start-service