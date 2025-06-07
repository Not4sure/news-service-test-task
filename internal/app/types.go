package app

import (
	"time"

	"github.com/not4sure/news-service-test-task/internal/domain/article"
)

type ArticleViewModel struct {
	ID        string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func viewModelFromDomainArticle(a article.Article) ArticleViewModel {
	return ArticleViewModel{
		ID:        a.ID().String(),
		Title:     a.Title(),
		Content:   a.Content(),
		CreatedAt: a.CreatedAt(),
		UpdatedAt: a.UpdatedAt(),
	}
}
