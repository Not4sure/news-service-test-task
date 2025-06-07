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

// New creates in-memory article Repository.
func New() article.Repository {
	return &Repository{
		articles: make(map[uuid.UUID]article.Article),
	}
}

// GetByID implements article.Repository.
func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (article.Article, error) {
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

	aa := []article.Article{}
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

// Update implements article.Repository.
func (r *Repository) Update(ctx context.Context, id uuid.UUID, update func(*article.Article) error) error {
	r.Lock()
	defer r.Unlock()

	a, ok := r.articles[id]
	if !ok {
		return article.ErrArticleNotFound
	}

	err := update(&a)
	if err != nil {
		return err
	}

	r.articles[id] = a
	return nil
}

// Delete implements article.Repository.
func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.articles[id]; !ok {
		return article.ErrArticleNotFound
	}

	delete(r.articles, id)
	return nil
}
