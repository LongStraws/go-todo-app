package main

import (
	"context"
	"database/sql"
	"fmt"
	"longstraws/go-todo-app/db"

	_ "modernc.org/sqlite"
)

var ddl string

func initDatabase(ctx context.Context, database *sql.DB) error {

	// init users
	if _, err := database.ExecContext(ctx, "CREATE TABLE users (userId INTEGER PRIMARY KEY, name text NOT NULL);"); err != nil {
		return err
	}

	// init projects
	if _, err := database.ExecContext(ctx, "create TABLE projects (projectId INTEGER PRIMARY KEY, name text NOT NULL);"); err != nil {
		return err
	}

	// init tasks
	if _, err := database.ExecContext(ctx, "CREATE TABLE tasks (id INTEGER PRIMARY KEY,name text NOT NULL,FOREIGN KEY (projectId) REFERENCES projects(projectId));"); err != nil {
		return err
	}

	return nil
}
func run() error {
	ctx := context.Background()

	database, err := sql.Open("sqlite", ":memory:")

	if err != nil {
		return err
	}

	initDatabase(ctx, database)

	// ExecContext executes a query without returning rows
	if _, err := database.ExecContext(ctx, ddl); err != nil {
		return err
	}

	queries := db.New(database)

	user1, err := queries.CreateUser(ctx, "test")

	if err != nil {
		return err
	}

	fmt.Println(user1.Userid)

	return nil
}

func main() {
	fmt.Println("test")

	if err := run(); err != nil {
		panic(err)
	}

	fmt.Print("Success!!")

}
