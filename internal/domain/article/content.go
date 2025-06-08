package article

import "errors"

const (
	MaxContentLen = 4096
)

var (
	ErrEmptyContent   = errors.New("content is empty")
	ErrContentTooLong = errors.New("content is too long")
)

type Content string

func NewContent(c string) (Content, error) {
	if len(c) == 0 {
		return "", ErrEmptyContent
	}
	if len(c) > MaxContentLen {
		return "", ErrContentTooLong
	}

	return Content(c), nil
}

func (c Content) String() string {
	return string(c)
}
