# Architecture Overview — Hello World Acceptance 5

## Stack

| Part | Choice |
|---|---|
| Frontend | Next.js 15 App Router, TypeScript, Tailwind v3, ESLint |
| Backend | Go 1.25-compatible module, `net/http`, pgx v5 |
| Database | PostgreSQL 16 |
| Runtime | `docker compose --profile local up --build` |

## Layout

```text
code/backend/cmd/api/main.go      API entry point and migration runner
code/backend/migrations/          Ordered SQL migration pairs
code/frontend/app/                App Router shell and frozen shared CSS
code/frontend/components/         One default-exported component per story
code/frontend/lib/mock/           Story mock data; delete when API lands
docs/architecture/                This overview, ERD, API contract
```

## Contracts and conventions

- `app/page.tsx` stays Server Component composition root. Story UI adds one component import and element.
- Interactive components start with literal first line `"use client"` and use `export default function ComponentName()`.
- CSS modules use only tokens from `app/globals.css`; globals defines approved color, spacing, typography, radius, shadow, motion tokens.
- Backend routes use `/v1/...`, never `/api`; proxy owns `/api` prefix.
- API returns JSON. Error body is `{ "error": { "code": "...", "message": "..." } }`.
- Backend reads `DATABASE_URL`, runs pending SQL migrations in filename order, then listens on `PORT`, `APP_PORT`, or `8080`.
- `GET /healthz` returns 200 only after migrations and database `SELECT 1` succeed.
- One shared greeting row. PostgreSQL transaction gives later completed valid write final state.

## Decisions

| Decision | Rejected | Tradeoff |
|---|---|---|
| One-row `greetings` table | In-memory value | Small schema gives restart persistence required by SRS. |
| pgx + handwritten SQL | ORM | Less abstraction and dependency surface; queries stay short. |
| Backend self-migrations | Manual migration command | Boot is reliable on empty runtime DB; startup waits for schema. |
| Client fetch for greeting UI | Server-only form | Client form needed for submit and focus behavior. |

## Environment

| Service | Keys |
|---|---|
| Backend | `DATABASE_URL`, `PORT`, optional `APP_PORT` fallback |
| Frontend | `NEXT_PUBLIC_API_URL`, `API_ORIGIN` |
| Compose | `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`; optional ports and resource limits |

Tracked `.env.example` files list keys with comments. Never commit values from local `.env` files.

## Run and verify

1. Copy root `.env.example` to `.env` if overrides are needed.
2. Run `docker compose --profile local up --build` from repository root.
3. Open `http://localhost:3000`; backend health endpoint is `http://localhost:8080/healthz`.

Migration files use sortable timestamp prefixes and paired `.up.sql`/`.down.sql`. Current initial migration is safe inside a transaction. Future `CREATE INDEX CONCURRENTLY` migration must run outside transaction; runner detects it. No rollout blocker: migration seeds initial greeting with conflict-safe insert.
