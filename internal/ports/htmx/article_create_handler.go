package htmx

import (
	"fmt"
	"net/http"
)

func (hs *HTMXServer) ArticleCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating article")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	title := r.Form.Get("title")
	content := r.Form.Get("content")

	a, err := hs.app.CreateArticle(r.Context(), title, content)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}

	err = hs.templates.ExecuteTemplate(w, "article-card.html", a)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
