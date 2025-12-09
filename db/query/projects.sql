-- name: ListProjects :many
SELECT * FROM projects;

-- name: CreateProject :one
INSERT INTO projects (
name
) values(
?
) RETURNING *;
