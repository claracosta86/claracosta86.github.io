package config

import (
    "context"
    "database/sql"
    "log"
    "os"
    "path/filepath"

    _ "github.com/mattn/go-sqlite3"

)

type DB struct {
    sql *sql.DB
}

type DBReader struct {
    db  *sql.DB
    ctx context.Context
}

type DBWriter struct {
    db  *sql.DB
    ctx context.Context
}

var dataBase *DB

func InitDB() *DB {
    dbFile := "./data.db"
    sqlDB, err := sql.Open("sqlite3", dbFile)
    if err != nil {
        log.Fatalf("Erro ao abrir o banco de dados: %v", err)
    }

    runInitSQL(sqlDB)

   return &DB{sql: sqlDB}
}

// Executa init.sql uma vez
func runInitSQL(db *sql.DB) {
    initSQLPath := filepath.Join("sql", "init.sql")
    sqlBytes, err := os.ReadFile(initSQLPath)
    if err != nil {
        log.Fatalf("Erro ao ler init.sql: %v", err)
    }

    _, err = db.Exec(string(sqlBytes))
    if err != nil {
        log.Fatalf("Erro ao executar init.sql: %v", err)
    }
}

// Leitura (SELECT)
func (db *DB) Read(ctx context.Context) *DBReader {
    return &DBReader{db: db.sql, ctx: ctx}
}

// Escrita (INSERT, UPDATE, DELETE)
func (db *DB) Write(ctx context.Context) *DBWriter {
    return &DBWriter{db: db.sql, ctx: ctx}
}

// Executa SELECT e retorna *sql.Rows
func (r *DBReader) Select(query string, args ...any) (*sql.Rows, error) {
    return r.db.QueryContext(r.ctx, query, args...)
}

// Executa comandos de escrita
func (w *DBWriter) Exec(query string, args ...any) (sql.Result, error) {
    return w.db.ExecContext(w.ctx, query, args...)
}
