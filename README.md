# Movie CRUD API in Go (No DB)

This project implements a simple CRUD (Create, Read, Update, Delete) API for managing movies using Go. Instead of using an actual database, it uses an in-memory struct to store movie data.

## Requirements

- Go 1.15 or higher

## Getting Started

1. Clone the repository:

```
   git clone https://github.com/your-username/movie-crud-api.git
   cd movie-crud-api
```

2. Install dependencies:

```
go mod tidy
```

3. Run the server:

```
go run main.go
```

The server should now be running at http://localhost:8080.

# Endpoints

## Get All Movies

- URL: /movies
- Method: GET
- Response: List of all movies in JSON format.

## Get a Single Movie

- URL: /movies/{id}
- Method: GET
- Response: Movie with the specified ID in JSON format.

## Create a New Movie

- URL: /movies
- Method: POST
- Request Body: Movie data in JSON format (ID, Title, Director).
- Response: Newly created movie in JSON format.

## Update a Movie

- URL: /movies/{id}
- Method: PUT
- Request Body: Updated movie data in JSON format (ID, Title, Director).
- Response: Updated movie in JSON format.

## Delete a Movie

- URL: /movies/{id}
- Method: DELETE
- Response: Deleted movie in JSON format.

## Movie Structure

A movie is represented by the following JSON structure:

```
{
  "ID": "string",
  "Title": "string",
  "Director": {
    "Firstname": "string",
    "Lastname": "string"
  }
}

```

## Sample Requests

1. Get All Movies:

```
GET /movies
```

2. Get a Single Movie

```
GET /movies/{id}

```

3. Create a New Movie:

```
POST /movies
Content-Type: application/json

{
  "ID": "movie_123",
  "Title": "New Movie Title",
  "Director": {
    "Firstname": "John",
    "Lastname": "Doe"
  }
}

```

4. Updated A Movie

```
PUT /movies/{id}
Content-Type: application/json

{
  "ID": "movie_123",
  "Title": "Updated Movie Title",
  "Director": {
    "Firstname": "Jane",
    "Lastname": "Smith"
  }
}

```

5. Delete a Movie:

```
DELETE /movies/{id}

```

## Error Handling

- If a movie with the given ID is not found, the server will respond with a 404 Not Found error.
- For other errors, the server will respond with a 500 Internal Server Error.

## License

This project is licensed under the MIT License. Feel free to use, modify, and distribute it as per the terms of the license.
