# Movie CRUD API

This is a simple REST API for performing CRUD operations on a collection of movies. It allows you to get all movies, get a specific movie by ID, create a new movie, update an existing movie, and delete a movie.

## Endpoints

- `GET /movies` - Retrieve all movies
- `GET /movies/{id}` - Retrieve a specific movie by ID
- `POST /movies` - Create a new movie
- `PUT /movies/{id}` - Update an existing movie by ID
- `DELETE /movies/{id}` - Delete a movie by ID

## Tech Stack

- **Language**: Go
- **Framework**: Gorilla Mux for routing

## Running the Application

1. Ensure you have Go installed.
2. Run `go run main.go` to start the server on port 8080.
