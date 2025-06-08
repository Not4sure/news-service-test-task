package article

import "errors"

const (
	MaxTitleLen = 256
)

var (
	ErrEmptyTitle   = errors.New("title is empty")
	ErrTitleTooLong = errors.New("title is too long")
)

type Title string

func NewTitle(title string) (Title, error) {
	if len(title) == 0 {
		return "", ErrEmptyTitle
	}
	if len(title) > MaxTitleLen {
		return "", ErrTitleTooLong
	}

	return Title(title), nil
}

func (t Title) String() string {
	return string(t)
}
