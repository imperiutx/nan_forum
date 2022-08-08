package db

import (
	"context"
	"github.com/imperiutx/nan_forum/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateTopic(t *testing.T) {
	arg := CreateTopicParams{
		CategoryID: 1,
		Title:      utils.RandomString(15),
		Body:       utils.RandomString(100) + "\n" + utils.RandomString(50),
		CreatedBy:  utils.RandomString(10),
	}

	topic, err := testQueries.CreateTopic(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, topic)

	require.Equal(t, arg.Title, topic.Title)
	require.Equal(t, arg.CreatedBy, topic.CreatedBy)

	require.NotZero(t, topic.ID)
	require.NotZero(t, topic.CreatedAt)
}

func TestGetTopicByID(t *testing.T) {
	topic, err := testQueries.GetTopicByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, topic)

	require.Equal(t, topic.Title, "pqsopfphkavnuis")
	require.Equal(t, topic.CreatedBy, "gtmtybpxwa")
}

func TestListTopics(t *testing.T) {
	topics, err := testQueries.ListTopics(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, topics)

	require.Equal(t, len(topics), 3)
	require.Equal(t, topics[2].CreatedBy, "msjrwhchby")
}

func TestUpdateTopicByID(t *testing.T) {

	arg := UpdateTopicByIDParams{
		ID:    2,
		Title: "Updated topic name 4",
		Body:  "updated contains again" + "\n" + utils.RandomString(100),
	}
	err := testQueries.UpdateTopicByID(context.Background(), arg)
	require.NoError(t, err)

	topic, err := testQueries.GetTopicByID(context.Background(), 2)
	require.NoError(t, err)
	require.NotEmpty(t, topic)

	require.Equal(t, topic.Title, "Updated topic name 4")

}

func TestHideTopicByID(t *testing.T) {
	err := testQueries.HideTopicByID(context.Background(), 3)
	require.NoError(t, err)

	topic, err := testQueries.GetTopicByID(context.Background(), 3)
	require.Error(t, err)
	require.Empty(t, topic)
}
