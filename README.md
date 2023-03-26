Sure, here's a sample README.md file:

# Go API

[![Go Report Card](https://goreportcard.com/badge/github.com/username/repo)](https://goreportcard.com/report/github.com/username/repo)

A simple API built with Go.

## Dependencies

- Go 1.16 or later

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/username/repo.git
   ```

2. Install the dependencies:

   ```
   go mod download
   ```

## Usage

To run the application:

```
go run main.go router.go controller.go todo.go
```

The application will be available at `http://localhost:8080`.

## Docker

To run the application using Docker:

1. Build the Docker image:

   ```
   docker build -t go-api .
   ```

2. Run the Docker container:

   ```
   docker run -p 8080:8080 go-api
   ```

## Files

- `main.go`: The main entry point of the application.
- `router.go`: Defines the routes for the API using the Gorilla Mux router.
- `controller.go`: Defines the controller functions for the API.
- `todo.go`: Defines the Todo struct used by the API.
- `Dockerfile`: Defines the Docker image for the application.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.