package htmx

import (
	"fmt"
	"net/http"

	"github.com/not4sure/news-service-test-task/internal/app"
)

func (hs *HTMXServer) ArticleCreateFormHandler(w http.ResponseWriter, r *http.Request) {
	err := hs.templates.ExecuteTemplate(w, "article-create-form.html", app.ArticleViewModel{})

	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
