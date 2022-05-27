package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	arg := CreateCategoryParams{
		Name:      "All categories",
		CreatedBy: "admin",
	}
	category, err := testQueries.CreateCategory(
		context.Background(),
		arg,
	)

	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Name, category.Name)
	require.Equal(t, arg.CreatedBy, category.CreatedBy)

	require.NotZero(t, category.ID)
	require.NotZero(t, category.CreatedAt)
}

func TestGetCategoryByID(t *testing.T) {
	category, err := testQueries.GetCategoryByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, category.Name, "All Categories")
	require.Equal(t, category.CreatedBy, "admin")
}

func TestListCategories(t *testing.T) {
	categories, err := testQueries.ListCategories(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, categories)

	require.Equal(t, categories[0].Name, "All Categories")
	require.Equal(t, categories[0].CreatedBy, "admin")
}
