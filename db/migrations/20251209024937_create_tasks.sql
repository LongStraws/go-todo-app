-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
id INTEGER PRIMARY KEY,
name text NOT NULL,
projectId INTEGER NOT NULL,
FOREIGN KEY (projectId) REFERENCES projects(projectId)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
