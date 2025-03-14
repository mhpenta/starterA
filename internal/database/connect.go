package database

import (
	"context"
	"database/sql"
	"fmt"
	"starterA/internal/config"
	"starterA/internal/database/repo"
	"time"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

// GetConnection establishes a connection to the Turso database
// and returns a database handle with proper connection pool settings
func GetConnection(dbCfg config.Database) (*sql.DB, error) {
	db, err := sql.Open("libsql", dbCfg.TursoConnectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		err := db.Close()
		if err != nil {
			return nil, fmt.Errorf("error closing database connection: %w", err)
		}
		return nil, err
	}

	return db, nil
}

// GetDatabase initializes a database connection and returns
// a Queries object that provides type-safe access to all database operations
func GetDatabase(dbCfg config.Database) (*repo.Queries, error) {
	dbConn, err := GetConnection(dbCfg)
	if err != nil {
		return nil, fmt.Errorf("error getting db connection: %w", err)
	}

	return repo.New(dbConn), nil
}
