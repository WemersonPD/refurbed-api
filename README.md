# Backend

## Architecture

```
┌─────────────┐   ┌──────────────┐   ┌───────────────┐
│  Controller  │──▶│   Service    │──▶│  Repository   │
│ (HTTP layer) │   │  (+ Cache)  │   │ (data access) │
└─────────────┘   └──────────────┘   └───────┬───────┘
                    In-memory                 │
                   cache (30s)                ▼
                                      ┌───────────────┐
                                      │  JSON Files   │
                                      │ metadata.json │
                                      │ details.json  │
                                      └───────────────┘
```

Each layer depends on interfaces, enabling dependency injection and testability.

## Prerequisites

- Go 1.22 or higher

## Running the Server

```bash
go run .
# Server starts on http://localhost:8080
```

## Running Tests

```bash
go test ./...        # all tests
go test ./... -v     # verbose output
```

## API Documentation

Full API documentation with examples is available on Postman:
https://documenter.getpostman.com/view/38503833/2sBXigMtLD

## Task Management

Project tasks were tracked via Trello board:
https://trello.com/invite/b/69b54bfdd5c96c11f3d54219/ATTIb94613ef684c7d4a1ce2c93692539bd67418FB4F/refurbed

## Data Files

- `data/metadata.json` - Product metadata (id, name, base_price, image_url)
- `data/details.json` - Product details (id, discount_percent, bestseller, colors, stock, category, brand, condition)

## Production Improvements

- **Database:** Replace JSON files with PostgreSQL for indexing, joins, and query performance.
- **Cache:** Use Redis for horizontal scalability and cache sharing across instances.
- **Observability:** Add structured logging, metrics (Prometheus), and tracing (OpenTelemetry).
- **Input validation:** Stricter validation and sanitization on query parameters.
- **Error handling:** Structured error codes with centralized error handling.
- **CI/CD:** Linting, test coverage thresholds, and automated deployments.
- **Rate limiting:** Protect the API from abuse.
- **Docker:** Containerize for consistent environments.

## Final Thoughts

The architecture is simple but extensible — swapping JSON files for a database or the in-memory cache for Redis would require changes only in the repository and cache layers, leaving the service and controller untouched.
