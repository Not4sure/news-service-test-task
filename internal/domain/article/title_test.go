package article_test

import (
	"math/rand"
	"testing"

	"github.com/not4sure/news-service-test-task/internal/domain/article"
	"github.com/stretchr/testify/require"
)

func TestCreateTitle(t *testing.T) {
	testCases := []struct {
		name  string
		title string
		err   error
	}{
		{
			name:  "OK min",
			title: stringOfLen(1),
			err:   nil,
		},
		{
			name:  "OK max",
			title: stringOfLen(article.MaxTitleLen),
			err:   nil,
		},
		{
			name:  "Empty title",
			title: "",
			err:   article.ErrEmptyTitle,
		},
		{
			name:  "Too long title",
			title: stringOfLen(article.MaxTitleLen + 1),
			err:   article.ErrTitleTooLong,
		},
	}

	t.Parallel()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			title, err := article.NewTitle(tc.title)

			require.Equal(t, tc.err, err)
			if err == nil {
				require.Equal(t, tc.title, title.String())
			}
		})
	}
}

func stringOfLen(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
