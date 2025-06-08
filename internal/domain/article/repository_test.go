package article_test

import (
	"context"
	"testing"

	"github.com/not4sure/news-service-test-task/internal/domain/article"
	"github.com/not4sure/news-service-test-task/internal/domain/article/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository(t *testing.T) {

	repositories := createRepositories()

	for i := range repositories {
		r := repositories[i]

		t.Run(r.Name, func(t *testing.T) {
			t.Parallel()

			t.Run("testStoreArticle", func(t *testing.T) {
				t.Parallel()
				testStoreArticle(t, r.Repository)
			})
		})
	}
}

// testStoreArticle tests article.Repository Store() func.
func testStoreArticle(t *testing.T, repo article.Repository) {
	ctx := context.Background()

	testCases := []struct {
		name          string
		createArticle func(t *testing.T) *article.Article
	}{
		{
			name: "valid_article",
			createArticle: func(t *testing.T) *article.Article {
				return newValidArticle(t)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := tc.createArticle(t)

			err := repo.Store(ctx, *a)
			require.NoError(t, err)

			assertArticleInRepository(ctx, t, repo, a)
		})
	}
}

func assertArticleInRepository(ctx context.Context, t *testing.T, repo article.Repository, a *article.Article) {
	articleFromRepo, err := repo.GetByID(ctx, a.ID())
	require.NoError(t, err)

	assert.Equal(t, a.ID(), articleFromRepo.ID())
	assert.Equal(t, a.Title(), articleFromRepo.Title())
	assert.Equal(t, a.Content(), articleFromRepo.Content())
	assert.Equal(t, a.CreatedAt(), articleFromRepo.CreatedAt())
	assert.Equal(t, a.UpdatedAt(), articleFromRepo.UpdatedAt())
}

// newValidArticle creates valid article with random values.
func newValidArticle(t *testing.T) *article.Article {
	a, err := article.NewArticle(stringOfLen(20), stringOfLen(1000))
	require.NoError(t, err)

	return a
}

type repo struct {
	Name       string
	Repository article.Repository
}

func createRepositories() []repo {
	return []repo{
		{
			Name:       "Memory",
			Repository: memory.New(),
		},
	}
}
