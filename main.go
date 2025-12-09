package main

import (
	"context"
	"database/sql"
	"fmt"
	"longstraws/go-todo-app/cmd/todo-app/root"
	"longstraws/go-todo-app/db"

	_ "embed"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

var ddl string

func initDb(ctx context.Context) (context.Context, *sql.DB, error) {
	// Initialize database connection here

	dbConn, err := sql.Open("sqlite", "testDb.sqlite")
	if err != nil {
		return nil, nil, err
	}

	if err := goose.Up(dbConn, "db/migrations"); err != nil {
		fmt.Println("Error migrating database:", err)
		return nil, nil, err
	}

	// create tables
	if _, err := dbConn.ExecContext(ctx, ddl); err != nil {
		return nil, nil, err
	}
	return ctx, dbConn, nil
}

func main() {
	ctx := context.Background()
	ctx, dbConn, err := initDb(ctx)

	queries := db.New(dbConn)
	if err != nil {
		fmt.Println(err)
		return
	}

	rootCmd := root.RootCommand(ctx, queries)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
