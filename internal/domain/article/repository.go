package article

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrArticleNotFound = errors.New("article not found")
)

type Repository interface {
	GetByID(ctx context.Context, id uuid.UUID) (Article, error)
	List(ctx context.Context) ([]Article, error)
	Store(ctx context.Context, a Article) error
	Update(ctx context.Context, id uuid.UUID, update func(*Article) error) error
	Delete(ctx context.Context, id uuid.UUID) error
}
