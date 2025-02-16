# README.md

# Go REST API

This project is a simple REST API built using Go and the Gin framework. It provides endpoints for managing albums and tracks HTTP request status codes using Prometheus metrics.

## Project Structure

```
go-rest-api
├── src
│   ├── main.go          # Entry point of the application
│   ├── albums
│   │   └── albums.go    # Logic for handling album-related requests
│   └── metrics
│       └── metrics.go   # Prometheus metrics setup
├── go.mod               # Module definition and dependencies
└── README.md            # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd go-rest-api
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run src/main.go
   ```

4. **Access the API:**
   - Get all albums: `GET http://localhost:3000/albums`
   - Create a new album: `POST http://localhost:3000/albums`

5. **Prometheus Metrics:**
   - The application exposes metrics at the `/metrics` endpoint for Prometheus to scrape.

## Usage

This API allows you to manage albums with basic CRUD operations. You can extend the functionality by adding more endpoints and metrics as needed.