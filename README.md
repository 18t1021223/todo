# Todo List Roadmap API

![Project Status](https://img.shields.io/badge/status-active-success.svg)
![Go Version](https://img.shields.io/badge/go-1.25-blue.svg)

**Project URL**: [https://github.com/18t1021223/todo)

## 🚀 Features

- **RESTful API**: Clean and consistent API design.
- **Authentication**: Secure JWT-based authentication (Access & Refresh tokens).
- **Database**: MySQL integration using `sqlc` for type-safe SQL generation.
- **Migrations**: Database schema management with `goose`.
- **Configuration**: Environment-based config using `viper`.
- **Logging**: Structured logging with `zap`.
- **Validation**: Request validation using `go-playground/validator`.
- **Pagination**: built-in support for paginated responses.

## 🛠 Tech Stack

- **Language**: [Go 1.25](https://golang.org/)
- **Router**: [Chi v5](https://github.com/go-chi/chi) - Lightweight, idiomatic and composable router.
- **Database**: MySQL
- **ORM/Query Builder**: [SQLC](https://sqlc.dev/) - Compile SQL to type-safe Go.
- **Migrations**: [Goose](https://github.com/pressly/goose)
- **Config**: [Viper](https://github.com/spf13/viper)
- **Logging**: [Zap](https://github.com/uber-go/zap)

## 📋 Prerequisites

Before you begin, ensure you have met the following requirements:

- **Go**: Version 1.25 or later.
- **MySQL**: A running MySQL database instance.
- **Make**: (Optional) For running convenience commands.

## ⚙️ Setup & Installation

1.  **Clone the repository**
    ```bash
    git clone <project_url>
    cd todo_list_roadmap
    ```

2.  **Configure Environment Variables**
    Copy the example environment file and update it with your credentials:
    ```bash
    cp .env.example .env
    ```
    Open `.env` and set your database credentials (`DB_HOST`, `DB_USER`, `DB_PASS`, `DB_NAME`) and JWT secret.

3.  **Run Database Migrations**
    Initialize the database schema:
    ```bash
    make goose-up
    ```
    *Alternatively, if you don't have `make` or `goose` installed globally, you may need to install goose or run the migration command manually.*

4.  **Install Dependencies**
    ```bash
    go mod download
    ```

## 🏃 Running the Application

You can start the server using the provided Makefile command:

```bash
make run
```

Or using standard Go command:

```bash
go run main.go
```

The server will start on port `8050` (or the port defined in `SERVER_PORT` in your `.env`).

## 🔌 API Endpoints

### User (Public)
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/register` | Register a new user account |
| `POST` | `/login` | Authenticate and receive a JWT token |

### Todos (Protected)
*Requires `Authorization: Bearer <token>` header*

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/todos` | Get a list of todos (supports pagination) |
| `POST` | `/todos` | Create a new todo item |
| `PUT` | `/todos/{id}` | Update an existing todo |
| `DELETE` | `/todos/{id}` | Delete a todo item |

## 📂 Project Structure

```
todo_list_roadmap/
├── config/         # Configuration initialization (DB, Env, Logger, JWT)
├── db/             # Database related files
│   ├── generated/  # SQLC generated Go code
│   ├── migrations/ # SQL migrations files
│   └── query/      # SQL queries for SQLC
├── dto/            # Data Transfer Objects (Request/Response models)
├── filter/         # Middleware (CORS, Auth, Logging)
├── handle/         # HTTP Handlers (Controllers)
├── router/         # Router setup and wiring
├── service/        # Business logic layer
├── util/           # Utility functions
├── main.go         # Application entry point
└── Makefile        # Build and run commands
```