# Use the official Go image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8080

# Start the application
CMD ["./main"]