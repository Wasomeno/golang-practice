-- name: CreateTodo :one

INSERT INTO todos (id, title, description, created_at, updated_at)
VALUES ($1,$2,$3,$4,$5)

RETURNING *;


-- name: GetTodo :one

SELECT * FROM todos
WHERE id = $1; 
