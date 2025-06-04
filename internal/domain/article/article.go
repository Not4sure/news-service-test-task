package article

import (
	"time"

	"github.com/google/uuid"
)

// Article represents news article in the system
type Article struct {
	id        uuid.UUID
	title     string
	content   string
	createdAt time.Time
	updatedAt time.Time
}

func NewArticle(
	title string,
	content string,
) *Article {
	// TODO: add title and content validation
	return &Article{
		id:        uuid.New(),
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}
}

func (a *Article) SetTitle(title string) error {
	a.title = title

	return nil
}

func (a *Article) SetContent(content string) error {
	a.content = content

	return nil
}

func (a Article) ID() uuid.UUID {
	return a.id
}

func (a Article) Title() string {
	return a.title
}

func (a Article) Content() string {
	return a.content
}

func (a Article) CreatedAt() time.Time {
	return a.createdAt
}

func (a Article) UpdatedAt() time.Time {
	return a.updatedAt
}
