# ERD — Greeting

## `greetings`

One shared current greeting. Application reads and updates row `id = 1`.

| Column | PostgreSQL type | Constraints | Purpose |
|---|---|---|---|
| `id` | `smallint` | Primary key; `CHECK (id = 1)` | Singleton key |
| `text` | `text` | `NOT NULL`; `CHECK (length(btrim(text)) > 0)` | Trimmed greeting |
| `updated_at` | `timestamptz` | `NOT NULL DEFAULT now()` | Last successful save |

Relationships: none. This plan needs one table only.

Initial migration inserts `(1, 'Hello, World!')` with `ON CONFLICT (id) DO NOTHING`, preserving prior saved state on restart or reapply.

## Story extension — Persisted editable greeting

No schema change beyond `greetings` required. This story is sole owner of singleton row and its current text. No foreign keys or indexes: reads and updates use primary key `id = 1`; no secondary query exists.

### UI mock review

`code/frontend/lib/mock/persisted-editable-greeting.ts` returns `GreetingResponse` as `{ "greeting": string }` for both read and save. Shape is sound and matches contracts below. Its localStorage persistence is mock-only; backend replaces module without changing component response handling. Mock silently returns current value for blank input; API instead returns `400 VALIDATION_ERROR`, required for boundary validation. Frontend keeps existing focus behavior and must handle that response without adding visible error UI.

## Migration plan

**Forward:** create `greetings` with `id smallint PRIMARY KEY CHECK (id = 1)`, `text text NOT NULL CHECK (length(btrim(text)) > 0)`, and `updated_at timestamptz NOT NULL DEFAULT now()`; then insert `(1, 'Hello, World!') ON CONFLICT (id) DO NOTHING`.

**Backward:** drop `greetings`. This removes current greeting and is only safe when rollback accepts data loss; no non-destructive reverse exists because table has no predecessor.

**Safety on populated databases:** safe. Creation is additive. Seed is idempotent and preserves row `id = 1`, so deploy/restart never overwrites a saved greeting. No table rewrite or lock beyond normal DDL transaction.
