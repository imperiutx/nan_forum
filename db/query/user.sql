-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 AND is_banned = false
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 AND is_banned = false
LIMIT 1;

-- name: GetUserByUserName :one
SELECT * FROM users
WHERE user_name = $1 AND is_banned = false
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
WHERE is_banned = false
ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (
    user_name, email, hashed_password, is_banned, is_admin
) VALUES (
             $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: BanUser :exec
UPDATE users
SET is_banned = true
WHERE id = $1;

-- name: UnBanUser :exec
UPDATE users
SET is_banned = false
WHERE id = $1;

-- name: SetAdmin :exec
UPDATE users
SET is_admin = true
WHERE id = $1;

-- name: UnSetAdmin :exec
UPDATE users
SET is_admin = false
WHERE id = $1;

-- name: UpdatePassword :exec
UPDATE users
SET hashed_password = $2
WHERE id = $1;