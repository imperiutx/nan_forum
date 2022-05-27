-- name: GetTopicByID :one
SELECT * FROM topics
WHERE id = $1 AND is_visible = true
LIMIT 1;

-- name: ListTopics :many
SELECT * FROM topics
WHERE is_visible = true
ORDER BY updated_at;

-- name: CreateTopic :one
INSERT INTO topics (
    category_id, title, body, created_by
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateTopicByID :exec
UPDATE topics
SET title = $2, body = $3, updated_at = now()
WHERE id = $1 AND is_visible = true;

-- name: HideTopicByID :exec
UPDATE topics
SET is_visible = false, updated_at = now()
WHERE id = $1;
