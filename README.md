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

## Environment Variables

The application uses environment variables for configuration. You can set these directly in your shell or use a `.env` file (which you'll need to load manually, e.g., using `source .env` or a tool like `direnv`).

**API Versioning:**
- `API_VERSION`: The version of the API, used in Swagger documentation (e.g., `1.0.0`).

**Database Configuration:**
- `DB_HOST`: PostgreSQL host (e.g., `localhost`)
- `DB_PORT`: PostgreSQL port (e.g., `5432`)
- `DB_USER`: PostgreSQL username
- `DB_PASSWORD`: PostgreSQL password
- `DB_NAME`: PostgreSQL database name

**Rate Limiting (Optional):**
- `RATE_LIMIT_COUNT`: Maximum number of requests allowed within the `RATE_LIMIT_WINDOW` (e.g., `1000`).
- `RATE_LIMIT_WINDOW`: Time duration for the rate limit window (e.g., `1m`, `30s`, `1h`). Uses Go's `time.ParseDuration` format.

**Example .env file:**
```
API_VERSION=1.0.0
DB_HOST=localhost
DB_PORT=5432
DB_USER=myuser
DB_PASSWORD=mypassword
DB_NAME=gahmen_db

RATE_LIMIT_COUNT=500
RATE_LIMIT_WINDOW=1m
```

## Running the Project

The project uses `Makefile` for common commands.

**Available `make` commands:**

- `make run`: Runs the API server.
- `make swag`: Generates the API documentation using Swaggo.
- `make test`: Runs all tests. (Assuming this command exists or will be added)

To run the project, ensure your environment variables are set, then use:

```bash
make run
```

To generate the API documentation:

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
