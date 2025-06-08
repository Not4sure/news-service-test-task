package mongo

import (
	"time"

	"github.com/google/uuid"
	"github.com/not4sure/news-service-test-task/internal/domain/article"
)

// articleSchema is a schema of article stored in a database
type articleModel struct {
	ID        uuid.UUID `bson:"_id"`
	Title     string    `bson:"title"`
	Content   string    `bson:"content"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func modelFromDomainArticle(a article.Article) articleModel {
	return articleModel{
		ID:        a.ID(),
		Title:     a.Title(),
		Content:   a.Content(),
		CreatedAt: a.CreatedAt(),
		UpdatedAt: a.UpdatedAt(),
	}
}

func domainArticleFromModel(m articleModel) (article.Article, error) {
	return article.UnmarhalFromDatabase(
		m.ID,
		m.Title,
		m.Content,
		m.CreatedAt,
		m.UpdatedAt,
	)
}
