package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ThanhNV121097/project-c5d18f08/backend/migrations"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("connect database: %v", err)
	}
	defer pool.Close()
	if err := migrate(ctx, pool); err != nil {
		log.Fatalf("apply migrations: %v", err)
	}
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("check database: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		if err := pool.Ping(r.Context()); err != nil {
			writeError(w, http.StatusServiceUnavailable, "SERVICE_UNAVAILABLE", "service unavailable")
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeError(w, http.StatusNotFound, "NOT_FOUND", "not found")
	})
	server := &http.Server{Addr: ":" + port(), Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	log.Printf("listening on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func migrate(ctx context.Context, pool *pgxpool.Pool) error {
	if _, err := pool.Exec(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (filename text PRIMARY KEY, applied_at timestamptz NOT NULL DEFAULT now())`); err != nil {
		return err
	}
	entries, err := fs.ReadDir(migrations.Files, ".")
	if err != nil { return err }
	var names []string
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".up.sql") { names = append(names, entry.Name()) }
	}
	sort.Strings(names)
	for _, name := range names {
		var applied bool
		if err := pool.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE filename = $1)`, name).Scan(&applied); err != nil { return err }
		if applied { continue }
		sql, err := migrations.Files.ReadFile(name)
		if err != nil { return err }
		if strings.Contains(strings.ToUpper(string(sql)), "CREATE INDEX CONCURRENTLY") {
			if _, err = pool.Exec(ctx, string(sql)); err != nil { return fmt.Errorf("%s: %w", name, err) }
			if _, err = pool.Exec(ctx, `INSERT INTO schema_migrations (filename) VALUES ($1)`, name); err != nil { return err }
			continue
		}
		tx, err := pool.Begin(ctx)
		if err != nil { return err }
		if _, err = tx.Exec(ctx, string(sql)); err == nil { _, err = tx.Exec(ctx, `INSERT INTO schema_migrations (filename) VALUES ($1)`, name) }
		if err != nil { _ = tx.Rollback(ctx); return fmt.Errorf("%s: %w", name, err) }
		if err = tx.Commit(ctx); err != nil { return err }
	}
	return nil
}

func port() string {
	if value := os.Getenv("PORT"); value != "" { return value }
	if value := os.Getenv("APP_PORT"); value != "" { return value }
	return "8080"
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil { log.Printf("write response: %v", err) }
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]map[string]string{"error": {"code": code, "message": message}})
}

var _ = errors.New
