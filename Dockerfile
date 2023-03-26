Sure! Here's the Dockerfile code:

```
# Start with a base golang image
FROM golang:1.16-alpine

# Set the working directory
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the binary
RUN go build -o main .

# Expose the port the application will run on
EXPOSE 8080

# Start the application
CMD ["./main"]
```