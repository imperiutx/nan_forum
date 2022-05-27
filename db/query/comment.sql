-- name: GetCommentByID :one
SELECT * FROM comments
WHERE id = $1 AND is_visible = true
LIMIT 1;

-- name: ListCommentsByTopicID :many
SELECT * FROM comments
WHERE topic_id = $1 AND is_visible = true
ORDER BY created_at;

-- name: CreateComment :one
INSERT INTO comments (
    topic_id, body, created_by
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateCommentByID :exec
UPDATE comments
SET body = $2, updated_at = now()
WHERE id = $1 AND is_visible = true;

-- name: HideCommentByID :exec
UPDATE comments
SET is_visible = false, updated_at = now()
WHERE id = $1;

-- name: IncreaseCommentPointsByID :exec
UPDATE comments
SET points = points + 1, updated_at = now()
WHERE id = $1 AND is_visible = true;

-- name: DecreaseCommentPointsByID :exec
UPDATE comments
SET points = points - 1, updated_at = now()
WHERE id = $1 AND is_visible = true;
