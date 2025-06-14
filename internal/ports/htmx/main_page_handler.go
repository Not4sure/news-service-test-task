package htmx

import (
	"net/http"

	"github.com/not4sure/news-service-test-task/internal/app"
)

type articlesCollection struct {
	Articles []app.ArticleViewModel
}

func (hs *HTMXServer) HandleMainPage(w http.ResponseWriter, r *http.Request) {
	articles, err := hs.app.ListArticles(r.Context())
	if err != nil {
		hs.handleAppErr(err, w)
		return
	}

	err = hs.templates.ExecuteTemplate(w, "index.html", articlesCollection{Articles: articles})
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
