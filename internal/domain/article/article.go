package article

import (
	"time"

	"github.com/google/uuid"
)

// Article represents news article in the system
type Article struct {
	id        uuid.UUID
	title     Title
	content   Content
	createdAt time.Time
	updatedAt time.Time
}

func NewArticle(
	title string,
	content string,
) (*Article, error) {
	t, err := NewTitle(title)
	if err != nil {
		return nil, err
	}

	c, err := NewContent(content)
	if err != nil {
		return nil, err
	}

	return &Article{
		id:        uuid.New(),
		title:     t,
		content:   c,
		createdAt: time.Now(),
	}, nil
}

// MustNewArticle creates article or panics on error.
// Intended for usage in tests.
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

// UnmarhalFromDatabase reconstructs Article from fields.
// Don't use this func for creation of new entities.
func UnmarhalFromDatabase(
	id uuid.UUID,
	title string,
	content string,
	createdAt time.Time,
	updatedAt time.Time,
) (Article, error) {
	t, err := NewTitle(title)
	if err != nil {
		return Article{}, err
	}

	c, err := NewContent(content)
	if err != nil {
		return Article{}, err
	}

	return Article{
		id:        id,
		title:     t,
		content:   c,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

// SetTitle updates title of an article.
func (a *Article) SetTitle(title string) error {
	t, err := NewTitle(title)
	if err != nil {
		return err
	}

	a.title = t
	a.updatedAt = time.Now()

	return nil
}

// SetContent updates content of an article.
func (a *Article) SetContent(content string) error {
	c, err := NewContent(content)
	if err != nil {
		return err
	}

	a.content = c
	a.updatedAt = time.Now()

	return nil
}

// ID is unique for each article entity.
func (a Article) ID() uuid.UUID {
	return a.id
}

// Title of article.
func (a Article) Title() string {
	return a.title.String()
}

// Content of article.
func (a Article) Content() string {
	return a.content.String()
}

// CreatedAt is a creation date of article.
func (a Article) CreatedAt() time.Time {
	return a.createdAt
}

// UpdatedAt is a last modification date of article.
func (a Article) UpdatedAt() time.Time {
	return a.updatedAt
}
