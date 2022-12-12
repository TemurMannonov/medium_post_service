package postgres_test

import (
	"testing"

	"github.com/TemurMannonov/medium_post_service/storage/repo"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createPost(t *testing.T) *repo.Post {
	category := createCategory(t)

	p, err := strg.Post().Create(&repo.Post{
		Title:       faker.Sentence(),
		Description: faker.Sentence(),
		UserID:      1,
		CategoryID:  category.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, p)

	return p
}

func TestCreatePost(t *testing.T) {
	createPost(t)
}
