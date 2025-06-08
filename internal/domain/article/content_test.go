package article_test

import (
	"testing"

	"github.com/not4sure/news-service-test-task/internal/domain/article"
	"github.com/stretchr/testify/require"
)

func TestCreateContent(t *testing.T) {
	testCases := []struct {
		name    string
		content string
		err     error
	}{
		{
			name:    "OK min",
			content: stringOfLen(1),
			err:     nil,
		},
		{
			name:    "OK max",
			content: stringOfLen(article.MaxContentLen),
			err:     nil,
		},
		{
			name:    "Empty content",
			content: "",
			err:     article.ErrEmptyContent,
		},
		{
			name:    "Cotent too long",
			content: stringOfLen(article.MaxContentLen + 1),
			err:     article.ErrContentTooLong,
		},
	}

	t.Parallel()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			content, err := article.NewContent(tc.content)

			require.Equal(t, tc.err, err)
			if err == nil {
				require.Equal(t, tc.content, content.String())
			}
		})
	}
}
