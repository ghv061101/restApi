-- name: CreateUser :one
INSERT INTO users (name, dob) VALUES (:name, :dob) RETURNING id, name, dob;

-- name: GetUserByID :one
SELECT id, name, dob FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, dob FROM users ORDER BY id;

-- name: UpdateUser :one
UPDATE users SET name = :name, dob = :dob WHERE id = :id RETURNING id, name, dob;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
