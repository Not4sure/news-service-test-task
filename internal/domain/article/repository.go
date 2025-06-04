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
	ByID(ctx context.Context, id uuid.UUID) (Article, error)
	List(ctx context.Context) ([]Article, error)
	Store(ctx context.Context, a Article) error
}
