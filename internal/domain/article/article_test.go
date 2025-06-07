package article_test

import (
	"testing"
	"time"

	"github.com/not4sure/news-service-test-task/internal/domain/article"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewArticle(t *testing.T) {
	testCases := []struct {
		name    string
		title   string
		content string
		err     error
	}{
		{
			name:    "OK",
			title:   "Test article",
			content: "Test article content",
			err:     nil,
		},
	}

	t.Parallel()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a, err := article.NewArticle(tc.title, tc.content)

			assert.Equal(t, tc.title, a.Title())
			assert.Equal(t, tc.content, a.Content())
			assert.WithinDuration(t, time.Now(), a.CreatedAt(), time.Second)
			assert.Equal(t, time.Time{}, a.UpdatedAt())
			require.Equal(t, tc.err, err)
		})
	}
}

func TestSetTitle(t *testing.T) {
	testCases := []struct {
		name  string
		title string
		err   error
	}{
		{
			name:  "OK",
			title: "Test title",
			err:   nil,
		},
	}

	t.Parallel()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := article.MustNewArticle("Test", "Test content")

			err := a.SetTitle(tc.title)

			require.Equal(t, tc.err, err)
			// Verify result only if update was successful.
			if err == nil {
				assert.Equal(t, tc.title, a.Title())
				assert.WithinDuration(t, time.Now(), a.UpdatedAt(), time.Second)
			}
		})
	}
}

func TestSetContent(t *testing.T) {
	testCases := []struct {
		name    string
		content string
		err     error
	}{
		{
			name:    "OK",
			content: "Test article content",
			err:     nil,
		},
	}

	t.Parallel()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := article.MustNewArticle("Test", "Test content")

			err := a.SetContent(tc.content)

			require.Equal(t, tc.err, err)
			// Verify result only if update was successful.
			if err == nil {
				assert.Equal(t, tc.content, a.Content())
				assert.WithinDuration(t, time.Now(), a.UpdatedAt(), time.Second)
			}
		})
	}
}
