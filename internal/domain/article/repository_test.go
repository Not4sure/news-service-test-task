package article_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/not4sure/news-service-test-task/internal/domain/article"
	"github.com/not4sure/news-service-test-task/internal/domain/article/memory"
	"github.com/not4sure/news-service-test-task/internal/domain/article/mongo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

var (
	mongoURI string
)

// TestMain starts docker container with mongodb for
func TestMain(m *testing.M) {
	ctx := context.Background()

	terminate, err := startMongoContainer(ctx)
	defer terminate(ctx)
	if err != nil {
		log.Fatalf("cannot start mongodb conainer: %e", err)
	}

	m.Run()
}

func startMongoContainer(ctx context.Context) (func(context.Context) error, error) {
	dbContainer, err := mongodb.Run(ctx, "mongo:latest")
	if err != nil {
		return nil, err
	}

	terminate := func(ctx context.Context) error {
		return dbContainer.Terminate(ctx)
	}

	dbHost, err := dbContainer.Host(ctx)
	if err != nil {
		return terminate, err
	}

	dbPort, err := dbContainer.MappedPort(ctx, "27017/tcp")
	if err != nil {
		return terminate, err
	}

	mongoURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)

	return terminate, err
}

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

// assertArticleInRepository checks if article is stored in repository
func assertArticleInRepository(ctx context.Context, t *testing.T, repo article.Repository, a *article.Article) {
	articleFromRepo, err := repo.GetByID(ctx, a.ID())
	require.NoError(t, err)

	assert.Equal(t, a.ID(), articleFromRepo.ID())
	assert.Equal(t, a.Title(), articleFromRepo.Title())
	assert.Equal(t, a.Content(), articleFromRepo.Content())
	assert.WithinDuration(t, a.CreatedAt(), articleFromRepo.CreatedAt(), time.Second)
	assert.WithinDuration(t, a.UpdatedAt(), articleFromRepo.UpdatedAt(), time.Second)
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
		{
			Name:       "Mongodb",
			Repository: mustNewMongoRepository(),
		},
	}
}

func mustNewMongoRepository() article.Repository {
	ctx := context.Background()
	repo := mongo.New(ctx, mongoURI, "news")

	return repo
}
