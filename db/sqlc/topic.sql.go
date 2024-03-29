// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: topic.sql

package db

import (
	"context"
)

const createTopic = `-- name: CreateTopic :one
INSERT INTO topics (
    category_id, title, body, created_by
) VALUES (
    $1, $2, $3, $4
) RETURNING id, category_id, title, body, created_by, points, is_visible, created_at, updated_at
`

type CreateTopicParams struct {
	CategoryID int64  `json:"category_id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	CreatedBy  string `json:"created_by"`
}

func (q *Queries) CreateTopic(ctx context.Context, arg CreateTopicParams) (Topic, error) {
	row := q.db.QueryRowContext(ctx, createTopic,
		arg.CategoryID,
		arg.Title,
		arg.Body,
		arg.CreatedBy,
	)
	var i Topic
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Title,
		&i.Body,
		&i.CreatedBy,
		&i.Points,
		&i.IsVisible,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const decreaseTopicPointsByID = `-- name: DecreaseTopicPointsByID :exec
UPDATE topics
SET points = points - 1, updated_at = now()
WHERE id = $1 AND is_visible = true
`

func (q *Queries) DecreaseTopicPointsByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, decreaseTopicPointsByID, id)
	return err
}

const getTopicByID = `-- name: GetTopicByID :one
SELECT id, category_id, title, body, created_by, points, is_visible, created_at, updated_at FROM topics
WHERE id = $1 AND is_visible = true
LIMIT 1
`

func (q *Queries) GetTopicByID(ctx context.Context, id int32) (Topic, error) {
	row := q.db.QueryRowContext(ctx, getTopicByID, id)
	var i Topic
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Title,
		&i.Body,
		&i.CreatedBy,
		&i.Points,
		&i.IsVisible,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const hideTopicByID = `-- name: HideTopicByID :exec
UPDATE topics
SET is_visible = false, updated_at = now()
WHERE id = $1
`

func (q *Queries) HideTopicByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, hideTopicByID, id)
	return err
}

const increaseTopicPointsByID = `-- name: IncreaseTopicPointsByID :exec
UPDATE topics
SET points = points + 1, updated_at = now()
WHERE id = $1 AND is_visible = true
`

func (q *Queries) IncreaseTopicPointsByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, increaseTopicPointsByID, id)
	return err
}

const listTopics = `-- name: ListTopics :many
SELECT id, category_id, title, body, created_by, points, is_visible, created_at, updated_at FROM topics
WHERE is_visible = true
ORDER BY updated_at DESC
`

func (q *Queries) ListTopics(ctx context.Context) ([]Topic, error) {
	rows, err := q.db.QueryContext(ctx, listTopics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Topic{}
	for rows.Next() {
		var i Topic
		if err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.Title,
			&i.Body,
			&i.CreatedBy,
			&i.Points,
			&i.IsVisible,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTopicByID = `-- name: UpdateTopicByID :exec
UPDATE topics
SET title = $2, body = $3, updated_at = now()
WHERE id = $1 AND is_visible = true
`

type UpdateTopicByIDParams struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (q *Queries) UpdateTopicByID(ctx context.Context, arg UpdateTopicByIDParams) error {
	_, err := q.db.ExecContext(ctx, updateTopicByID, arg.ID, arg.Title, arg.Body)
	return err
}
