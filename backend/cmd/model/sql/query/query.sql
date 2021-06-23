-- name: GetUser :one
SELECT * FROM "user"
WHERE id = $1;