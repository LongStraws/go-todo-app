-- +goose Up
-- +goose StatementBegin
CREATE TABLE projects (
projectId INTEGER PRIMARY KEY,
name text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE projects;
-- +goose StatementEnd
