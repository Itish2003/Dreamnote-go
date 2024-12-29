# Start from the Go base image for building the application
FROM golang:latest AS build

# Set the working directory in the container
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .



# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

# Use a minimal base image for the final container
FROM alpine:latest

# Install necessary runtime dependencies, such as certificates
RUN apk --no-cache add ca-certificates

# Set the working directory in the final container
WORKDIR /root/

# Copy the built server binary from the build stage
COPY --from=build /app/server .

# Expose the application port
EXPOSE 8080

# Command to run the server
CMD ["./server"]
