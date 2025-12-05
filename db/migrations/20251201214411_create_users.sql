-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
userId INTEGER PRIMARY KEY,
name text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
