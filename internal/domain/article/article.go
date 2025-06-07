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
) (*Article, error) {
	// TODO: add title and content validation
	return &Article{
		id:        uuid.New(),
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

// MustNewArticle creates article or panics on error.
// Intended for use in testing.
func MustNewArticle(
	title string,
	content string,
) *Article {
	a, err := NewArticle(title, content)
	if err != nil {
		panic(err)
	}

	return a
}

// SetTitle updates title of an article.
func (a *Article) SetTitle(title string) error {
	a.title = title
	a.updatedAt = time.Now()

	return nil
}

// SetContent updates content of an article.
func (a *Article) SetContent(content string) error {
	a.content = content
	a.updatedAt = time.Now()

	return nil
}

// ID is unique for each article entity.
func (a Article) ID() uuid.UUID {
	return a.id
}

// Title of article.
func (a Article) Title() string {
	return a.title
}

// Content of article.
func (a Article) Content() string {
	return a.content
}

// CreatedAt is a creation date of article.
func (a Article) CreatedAt() time.Time {
	return a.createdAt
}

// UpdatedAt is a last modification date of article.
func (a Article) UpdatedAt() time.Time {
	return a.updatedAt
}
