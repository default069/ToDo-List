# ToDo API

A simple REST API for managing a to-do list, built with **Go**, **Gin**, **GORM** and **PostgreSQL**, with interactive documentation via **Swagger**.

## Features

- Create, read, update and delete tasks
- PostgreSQL storage with automatic table migration
- Interactive API docs (Swagger UI)

## Tech Stack

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [GORM](https://gorm.io/) + PostgreSQL driver - ORM
- [swaggo/swag](https://github.com/swaggo/swag), [gin-swagger](https://github.com/swaggo/gin-swagger) - Swagger docs

## Project Structure

```
.
├── cmd
│   ├── docs            # Generated Swagger files (do not edit by hand)
│   └── main.go         # Entry point, routes, Swagger setup
├── db
│   ├── database.go
│   └── postgres.go     # PostgreSQL connection and auto-migration
├── internal
│   └── handlers
│       └── todo.go     # HTTP handlers
├── models
│   └── models.go       # Data models
├── go.mod
└── go.sum
```

## Prerequisites

- Go 1.22 or newer
- PostgreSQL running locally

## Getting Started

### 1. Clone the repository

```bash
git clone <https://github.com/default069/ToDo-List.git>
cd ToDo
```

### 2. Create the database

```bash
psql -U postgres -c "CREATE DATABASE todo_db;"
```

### 3. Configure the connection

The connection string is defined in `db/postgres.go`:

```go
dns := "postgres://postgres:1111@localhost:5432/todo_db?sslmode=disable"
```

Change the user, password, host or database name to match your setup.

> **Note:** do not commit real credentials. Consider moving this value to an environment variable.

### 4. Install dependencies

```bash
go mod tidy
```

### 5. Run the server

```bash
go run cmd/main.go
```

The API is available at `http://localhost:8080`. The `todos` table is created automatically on the first run.

## API Endpoints

| Method | Endpoint      | Description                    |
|--------|---------------|--------------------------------|
| GET    | `/todos`      | Get all tasks                  |
| GET    | `/todos/{id}` | Get a task by ID               |
| POST   | `/todos`      | Create a new task              |
| PATCH  | `/todos/{id}` | Partially update a task        |
| DELETE | `/todos/{id}` | Delete a task                  |

### Task model

```json
{
  "id": 1,
  "title": "Buy milk",
  "description": "2 liters, low fat",
  "completed": false
}
```

### Examples

Create a task:

```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Buy milk", "description": "2 liters", "completed": false}'
```

Get all tasks:

```bash
curl http://localhost:8080/todos
```

Mark a task as completed (only the sent fields are changed):

```bash
curl -X PATCH http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"completed": true}'
```

Delete a task:

```bash
curl -X DELETE http://localhost:8080/todos/1
```

## Swagger Documentation

Once the server is running, open:

**http://localhost:8080/swagger/index.html**

There you can view every endpoint and try requests directly from the browser.

### Regenerating the docs

Swagger files are generated from the comments above `main()` and the handlers. After changing any annotation, run from the project root:

```bash
swag init -g cmd/main.go -o cmd/docs
```

If the `swag` command is not found, install it and add Go's bin directory to your `PATH`:

```bash
go install github.com/swaggo/swag/cmd/swag@v1.16.6
export PATH=$PATH:$(go env GOPATH)/bin
```

Then restart the server.

## License

This project is for learning purposes. Add a license of your choice if you plan to share it.