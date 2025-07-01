# Gahmen-API Backend

This is a Go-based API that serves as a backend for the Gahmen Project

## Features

- List all ministries
- Get ministry details by ID
- List budget documents by ministry
- List expenditure by ministry
- List top N personnel by ministry
- Get project expenditure by ID
- Get programme expenditure by ministry ID
- List SGDI links by ministry ID

## Technologies

- [Go](https://golang.org/) - The programming language used
- [PostgreSQL](https://www.postgresql.org/) - The database used
- [Swaggo](https://github.com/swaggo/swag) - The library used for generating API documentation

## Installation

Follow these steps to install the project:

1. Clone the repository: `git clone https://github.com/bearylogical/gahmen_api.git`
2. Navigate to the project directory: `cd gahmen_api`
3. Download the dependencies: `go mod download`
4. Install Swaggo: `go get -u github.com/swaggo/swag/cmd/swag`
5. Set up the database: 
   - Create a PostgreSQL database
   - Create a `.env` file in the root directory and add the following environment variables:
     ```
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=user
     DB_PASSWORD=password
     DB_NAME=gahmen
     ```

## Running the Project

To run the project, use the following command:

```bash
go run ./cmd/app/main.go
```

To generate the API documentation, run the following command:

```bash
make swag
```

## API Documentation

The API documentation is generated using Swaggo and can be accessed at `http://localhost:3080/swagger/index.html` when the server is running.

## Endpoints

Here are some of the available endpoints:

- `GET /api/v1/ministries`
- `GET /api/v1/ministries/{ministry_id}`
- `GET /api/v1/budget/{ministry_id}/documents`
- `GET /api/v1/budget`
- `GET /api/v1/projects/{project_id}`
- `GET /api/v1/budget/{ministry_id}/programmes`
- `GET /api/v1/budget/opts`
- `GET /api/v1/budget/{ministry_id}`
- `GET /api/v1/budget/{ministry_id}/projects`
- `GET /api/v1/sgdi/{ministry_id}/links`
- `GET /api/v1/personnel`
- `GET /api/v2/budget`
- `POST /api/v2/projects`
