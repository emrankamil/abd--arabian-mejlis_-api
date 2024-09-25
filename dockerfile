# Use the official Golang image
FROM golang:1.22

# Install libvips for image compression
RUN apt-get update --fix-missing && apt-get install -y --no-install-recommends \
    libvips-dev \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# List files for debugging
RUN ls -al /app

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin ./cmd

# Expose the port
EXPOSE 8080

# Run the Go app
ENTRYPOINT ["/app/bin"]
