# Use the official Golang base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application runs on
EXPOSE 8080

# Run data seeder
# RUN main seed-superuser
# RUN main seed-data

# Command to run the executable
CMD ["./main", "seed-superuser"]
CMD ["./main", "seed-data"]
CMD ["./main"]