[![pt-br](https://img.shields.io/badge/lang-pt--br-green.svg)](https://github.com/oagarian/moviedle/blob/develop/README.pt-br.md)

# Moviedle

## What is it?
Moviedle is a project inspired by the "dle" games like [LoLdle](https://loldle.net/) and [Pokédle](https://pokedle.net/), but with a focus on movies. The goal is to provide a fun and interactive experience for movie enthusiasts.

## Dependencies
To run the project, you need to have the following installed:
1. [GoLang](https://golang.org/)
2. [Docker](https://www.docker.com/)
3. [Migrate package](https://pkg.go.dev/github.com/golang-migrate/migrate/v4)

## How to run the project
1. Clone the repository:
    ```sh
    git clone https://github.com/oagarian/moviedle
    ```
2. Download the project dependencies:
    ```sh
    go mod tidy
    ```
3. Start the database using the `execute.sh` script in the root of the project:
    ```sh
    chmod +x ./execute.sh
    ./execute.sh -environment -database
    ```
4. Run the tests:
    ```sh
    go test ./...
    ```
5. Start the application:
    ```sh
    go run src/apps/api/main.go
    ```

To test the routes, access the endpoints described in the Routes section.

## Routes

### API Documentation
`/api/docs/index.html`  
Access to see the full API documentation.

### Movie Endpoints
- **Create/Add a Movie**  
  `POST /api/movies`  
  Adds a new movie to the database.

- **List All Movies**  
  `GET /api/movies`  
  Lists all movies cataloged in the database.

- **Get Movie Details**  
  `GET /api/movies/{id}`  
  Displays details of a specific movie by ID.

- **List Movies by Filters**  
  `GET /api/movies?{filters}`  
  Lists movies based on specified filters.

- **Edit a Movie**  
  `PUT /api/movies/{id}`  
  Edits the information of a specific movie by ID.

- **Delete a Movie**  
  `DELETE /api/movies/{id}`  
  Removes a specific movie by ID from the database.

## Contributing
Contributions are welcome! Feel free to open issues and pull requests in the repository.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---
