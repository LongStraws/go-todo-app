-- name: ListTasks :many
SELECT * FROM tasks

-- name: CreateTask :one
INSERT INTO tasks (
name,
projectId
) values(
?,
?
) RETURNING *;
