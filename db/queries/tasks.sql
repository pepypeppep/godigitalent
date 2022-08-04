-- name: GetToDoList :many
SELECT * FROM tasks WHERE is_done = 0;

-- name: GetDoneList :many
SELECT * FROM tasks WHERE is_done = 1;

-- name: TaskInsert :execresult
INSERT INTO tasks (description, assignee, deadline_at)
VALUES (?,?,?);

-- name: TaskGetById :one
SELECT *
FROM tasks
WHERE id = ?;

-- name: TaskUpdate :exec
UPDATE tasks
set description= ?,
    assignee= ?,
    deadline_at= ?
where id = ?;

-- name: TaskUpdateProgress :exec
UPDATE tasks
set is_done= ?
where id = ?;

-- name: TaskDelete :exec
DELETE FROM tasks WHERE id = ?;