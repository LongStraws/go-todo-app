# go-todo-app

Simple SQLite-backed CLI for creating and listing todo projects. The app uses [Cobra](https://github.com/spf13/cobra) for commands, [`sqlc`](https://sqlc.dev) for type-safe queries, and [Goose](https://github.com/pressly/goose) to keep migrations in sync.

## Database & migrations
- Migrations live in `db/migrations/`. Ensure the target database file exists (e.g. `touch testDb.sqlite`) or let SQLite create it automatically.
- Apply migrations: `goose -dir db/migrations sqlite3 ./testDb.sqlite up`
- Roll back the last migration: `goose -dir db/migrations sqlite3 ./testDb.sqlite down`
- Create a new migration template: `goose -dir db/migrations create <name> sql`

## Build & run
```sh
go build -o todo-app main.go   # build binary
./todo-app                     # run root command (prints welcome message)
```

## Commands
- `./todo-app` – default root command, currently just greets you.
- `./todo-app add <project-name>` – creates a new project record via SQLC/SQLite.
- `./todo-app list` – lists all stored projects.

Add more subcommands under `cmd/todo-app/` to extend the CLI while reusing the shared database initialization logic in `main.go`.
