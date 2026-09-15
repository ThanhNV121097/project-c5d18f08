# Service Contracts — Greeting

Base path is backend origin. Deployment proxy strips browser `/api` prefix before backend routing; paths below deliberately omit it.

## Shared errors

All errors return `application/json`:

```json
{"error":{"code":"VALIDATION_ERROR","message":"greeting must contain non-whitespace text"}}
```

`code` is stable machine value. `message` is safe plain text. No internal DB detail leaks.

## Endpoints

### `GET /v1/greeting`

Returns current shared greeting. If row is absent, service creates or returns initial `Hello, World!` value.

- Success: `200 OK`

```json
{"greeting":"Hello, World!"}
```

- Failure: `500 INTERNAL_ERROR` using shared envelope.

### `PUT /v1/greeting`

Replaces current greeting. Service trims surrounding whitespace before validation and persistence. Last completed request wins.

Request:

```json
{"greeting":"Hello, Pipeline!"}
```

Success: `200 OK`

```json
{"greeting":"Hello, Pipeline!"}
```

Errors:

- `400 VALIDATION_ERROR`: missing, non-string, empty, or whitespace-only `greeting`.
- `500 INTERNAL_ERROR`: database or unexpected server failure.

Unknown paths return `404 NOT_FOUND` using shared envelope. Unsupported methods return `405 METHOD_NOT_ALLOWED` using shared envelope.

## Health

`GET /healthz` is runtime health only, returns `200 OK` with `{"status":"ok"}` after migration and database connectivity succeed; otherwise `503 Service Unavailable`.
