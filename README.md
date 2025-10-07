# Go Event Management REST API

Lightweight REST API built with Go and Gin that provides user authentication (JWT) and event management endpoints (create, read, update, delete). This repository contains handlers, models, queries, and simple middleware to secure routes using JSON Web Tokens.

## Project structure

Key files and folders:

- `main.go` — application entrypoint and router setup
- `db/` — database initialization (`db.go`)
- `models/` — data models (`user.go`, `event.go`)
- `handlers/` — HTTP handlers for events and users (`event.go`, `user.go`)
- `routes/` — route definitions (`event.go`, `user.go`)
- `middlewares/` — authentication middleware (`auth.go`)
- `utils/` — helper utilities (password hashing, JWT helpers)
- `queries/queries.go` — SQL queries used by the app
- `api-test/` — ready-to-use HTTP request examples (VS Code REST Client)

## Features

- User sign up and login with password hashing
- JWT-based authentication middleware to protect endpoints
- CRUD endpoints for events
- Example HTTP requests in `api-test/` for quick manual testing

## Requirements

- Go 1.20+ (check `go.mod` for exact version)
- sqlite3 (file-based DB is provided as `api.db`)

## Quick start

1. Install dependencies (module-aware):

```bash
go mod download
```

2. Build and run:

```bash
go run main.go
```

The server will start on the port configured by `APP_PORT` (or `8080`).

## API endpoints

Authentication:

- `POST /api/v1/signup` — create user (body: `username`, `email`, `password`)
- `POST /api/v1/login` — login and return JWT (body: `email`, `password`)

Events (protected — require `Authorization: Bearer <token>`):

- `GET /api/v1/events` — list all events
- `GET /api/v1/events/:id` — get single event
- `POST /api/v1/events` — create new event
- `PUT /api/v1/events/:id` — update existing event
- `DELETE /api/v1/events/:id` — delete an event

Refer to the handler function names in `handlers/` for request/response shapes.

## Using the provided API tests

This repository includes `api-test/` which contains HTTP requests you can run with the VS Code REST Client extension or by copying the requests to curl.

Examples:

- Sign up: `api-test/sign-up.http`
- Login: `api-test/login.http` (copy the returned `token` and use it in `Authorization` header for events)
- Create an event: `api-test/create-a-event.http`

To use curl for the login flow and create event (example):

```bash
# Login and extract token (POSIX shell example)
# For Windows cmd, adjust quoting and extraction methods accordingly.
curl -s -X POST http://localhost:8080/api/v1/login -H "Content-Type: application/json" -d '{"email":"you@example.com","password":"yourpass"}'

# Use the returned token in the Authorization header:
curl -X POST http://localhost:8080/api/v1/events -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"title":"My Event","description":"Desc","date":"2025-10-07"}'
```

Note: The quick curl examples above are POSIX-style; on Windows `cmd.exe` use double quotes with escaped inner quotes if needed.

## Tests

There are no automated tests included by default. You can run the HTTP examples in `api-test/` to exercise the API.

## Development notes

- Database: this project uses an SQLite file `api.db` in the repo root. Back it up before destructive testing.
- JWT: Use a strong secret in production.
- Passwords: hashed in `utils/hashing.go` before storage.

## Next steps / Improvements

- Add unit and integration tests
- Add migrations instead of a static DB file
- Add role-based access control and better error handling

## License

This project doesn't include a license file. Add a license if you plan to publish this repository.

---
