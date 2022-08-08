package db

import (
	"context"
	"github.com/imperiutx/nan_forum/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateComment(t *testing.T) {
	arg := CreateCommentParams{
		TopicID:   5,
		Body:      utils.RandomString(10) + " " + utils.RandomString(10),
		CreatedBy: utils.RandomString(10),
	}

	comment, err := testQueries.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, arg.TopicID, comment.TopicID)
	require.Equal(t, arg.Body, comment.Body)
	require.Equal(t, arg.CreatedBy, comment.CreatedBy)

	require.NotZero(t, comment.ID)
	require.NotZero(t, comment.CreatedAt)
}

func TestGetCommentByID(t *testing.T) {
	comment, err := testQueries.GetCommentByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, comment.CreatedBy, "mjtlcqrsyz")
}

func TestListCommentByTopicIDs(t *testing.T) {
	comments, err := testQueries.ListCommentsByTopicID(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, comments)

	require.Equal(t, len(comments), 1)
	require.Equal(t, comments[0].CreatedBy, "mjtlcqrsyz")
}

func TestUpdateCommentByID(t *testing.T) {

	arg := UpdateCommentByIDParams{
		ID:   1,
		Body: "Updated contains again",
	}
	err := testQueries.UpdateCommentByID(context.Background(), arg)
	require.NoError(t, err)

	comment, err := testQueries.GetCommentByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, comment.Body, "Updated contains again")

}

func TestHideCommentByID(t *testing.T) {
	err := testQueries.HideCommentByID(context.Background(), 3)
	require.NoError(t, err)

	comment, err := testQueries.GetCommentByID(context.Background(), 3)
	require.Error(t, err)
	require.Empty(t, comment)
}
