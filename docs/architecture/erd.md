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
