package memory

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/not4sure/news-service-test-task/internal/domain/article"
)

// Repository is an in-memory implementation of article.Repository.
type Repository struct {
	sync.Mutex
	articles map[uuid.UUID]article.Article
}

// NewRepository creates in-memory article Repository.
func NewRepository() article.Repository {
	return &Repository{
		articles: make(map[uuid.UUID]article.Article),
	}
}

// ByID implements article.Repository.
func (r *Repository) ByID(ctx context.Context, id uuid.UUID) (article.Article, error) {
	r.Lock()
	defer r.Unlock()

	if a, ok := r.articles[id]; ok {
		return a, nil
	}

	return article.Article{}, article.ErrArticleNotFound
}

// List implements article.Repository.
func (r *Repository) List(ctx context.Context) ([]article.Article, error) {
	r.Lock()
	defer r.Unlock()

	aa := make([]article.Article, len(r.articles))
	for _, a := range r.articles {
		aa = append(aa, a)
	}

	return aa, nil
}

// Store implements article.Repository.
func (r *Repository) Store(ctx context.Context, a article.Article) error {
	r.Lock()
	defer r.Unlock()

	r.articles[a.ID()] = a
	return nil
}
