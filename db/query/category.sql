-- name: GetCategoryByID :one
SELECT * FROM categories
WHERE id = $1 AND is_visible = true
LIMIT 1;

-- name: ListCategories :many
SELECT * FROM categories
WHERE is_visible = true
ORDER BY id;

-- name: CreateCategory :one
INSERT INTO categories (
    name, created_by
) VALUES (
    $1, $2
)
RETURNING *;