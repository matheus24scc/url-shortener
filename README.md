# URL Shortener

A lightweight URL shortener service built with Go, Gin, Redis, and PostgreSQL.

## Overview

Generate short, memorable URLs from long ones. Cache hot links in Redis and persist the
mapping in PostgreSQL. Docker Compose spins up the full stack (api + Redis + PostgreSQL)
locally and the OpenAPI spec documents the REST surface.
## Current Status (verified — checkup 2026-08-18)

> **Heads-up — the documentation below over-promises the current code.**
> This repository is currently a **minimal working stub**, not the full
> Postgres + Redis service described in the sections that follow.

What actually exists and is verified working (`go build ./...`, `go vet ./...`,
`go test ./...` all green):

- `cmd/server/main.go` — a Gin server exposing **only** `GET /ping` (returns `{"message":"pong"}`).
- `internal/config` — env-driven config loader (`config.Load()`).
- `internal/model` — the `URL` domain struct (defined, not yet wired to any storage).

What is **NOT implemented** yet (referenced by this README but absent from the code):

- `internal/handler`, `internal/middleware`, `internal/repository`, `internal/service` — directories do not exist.
- PostgreSQL (pgx) and Redis (go-redis) integration — `go.mod` lists only `gin-gonic/gin` and `joho/godotenv`; there are **zero** imports of `pgx`, `go-redis`, or `postgres` anywhere in the `.go` files.
- The `POST /shorten` and `GET /:code` endpoints described in the API table.
- The `docs/` directory and the Swagger UI (`swag init` has not been run / is not committed).

The `docker-compose.yml` / `Dockerfile` describe the intended full stack, but running
them would currently start a service that only serves `/ping`. Treat the Postgres/Redis
sections as the roadmap, not the present state.


## Tech Stack

| Layer    | Technology              |
|----------|-------------------------|
| Language | Go 1.22                 |
| Web      | Gin                     |
| Database | PostgreSQL 15 (pgx v4)  |
| Cache    | Redis 7 (go-redis v8)   |
| Docs     | OpenAPI 3.0.3 + Swagger |

## Project Structure

```
url-shortener/
├── cmd/server/main.go         # Application entry point
├── internal/
│   ├── config/                # Env-driven config loader
│   ├── handler/               # Gin HTTP handlers
│   ├── middleware/            # Cross-cutting HTTP middleware
│   ├── model/                 # Domain types (e.g. URL)
│   ├── repository/            # PostgreSQL data access
│   └── service/               # Business logic
├── api/openapi.yaml           # OpenAPI 3.0.3 spec
├── docs/                      # Generated swagger assets
├── Dockerfile                 # Multi-stage build (golang:1.22 → alpine)
├── docker-compose.yml         # api + postgres + redis stack
├── go.mod                     # Go module definition
└── .env.example               # Sample environment variables
```

## Prerequisites

- Go 1.22+
- Docker + Docker Compose (recommended)
- PostgreSQL 15+ and Redis 7+ (if running outside Docker)

## Quick Start (Docker)

```bash
# 1. Copy and adjust environment
cp .env.example .env

# 2. Build and start the stack
docker-compose up --build

# API is now available at http://localhost:8080
```

The compose file starts three services:

- `app` — the Go service on `:8080`
- `postgres` — PostgreSQL 15 on `:5432`
- `redis` — Redis 7 on `:6379`

## Local Development (without Docker)

```bash
# 1. Start Postgres and Redis (or use docker-compose for those pieces only)
cp .env.example .env

# 2. Run the service
go mod download
go run ./cmd/server
```

## API

See `api/openapi.yaml` for the full schema. The Swagger UI can be served locally once
the docs assets are generated (`swag init` produces the bundle under `docs/`).

### Endpoints

| Method | Path       | Description                    |
|--------|------------|--------------------------------|
| GET    | `/ping`    | Liveness probe (returns pong)  |
| POST   | `/shorten` | Create a short URL             |
| GET    | `/:code`   | Redirect to the original URL   |

## Configuration

All settings come from environment variables (see `.env.example`):

| Variable        | Description                    | Default       |
|-----------------|--------------------------------|---------------|
| `SERVER_PORT`   | HTTP listen port               | `8080`        |
| `DB_HOST`       | PostgreSQL host                | `localhost`   |
| `DB_PORT`       | PostgreSQL port                | `5432`        |
| `DB_USER`       | PostgreSQL user                | `postgres`    |
| `DB_PASSWORD`   | PostgreSQL password            | `postgres`    |
| `DB_NAME`       | PostgreSQL database            | `urlshortener`|
| `REDIS_ADDR`    | Redis address (`host:port`)    | `localhost:6379` |
| `REDIS_PASSWORD`| Redis password (if any)        | _empty_       |
| `REDIS_DB`      | Redis logical DB index         | `0`           |

## Building the Binary

```bash
go build -o bin/url-shortener ./cmd/server
./bin/url-shortener
```

The included `Dockerfile` does this in a multi-stage build (`golang:1.22-alpine` → `alpine:latest`).

## License

MIT
