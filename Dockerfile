# Use the official Go image as the base image
FROM golang:1.22.4

# Set the working directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Tidy up the modules to ensure no missing or extraneous dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port on which the app will run
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
