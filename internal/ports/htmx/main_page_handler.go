package htmx

import (
	"fmt"
	"net/http"

	"github.com/not4sure/news-service-test-task/internal/app"
)

type articlesCollection struct {
	Articles []app.ArticleViewModel
}

func (hs *HTMXServer) HandleMainPage(w http.ResponseWriter, r *http.Request) {
	articles, err := hs.app.ListArticles(r.Context())
	if err != nil {
		http.Error(w, "failed to list articles", http.StatusInternalServerError)
	}

	err = hs.templates.ExecuteTemplate(w, "index.html", articlesCollection{Articles: articles})

	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
