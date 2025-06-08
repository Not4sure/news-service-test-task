package htmx

import (
	"fmt"
	"net/http"

	"github.com/not4sure/news-service-test-task/internal/app"
)

func (hs *HTMXServer) ArticleCreatorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create article form")

	err := hs.templates.ExecuteTemplate(w, "article-creator.html", app.ArticleViewModel{})

	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
