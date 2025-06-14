package htmx

import (
	"net/http"
)

func (hs *HTMXServer) ArticleCreateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	title := r.Form.Get("title")
	content := r.Form.Get("content")

	a, err := hs.app.CreateArticle(r.Context(), title, content)
	if err != nil {
		hs.handleAppErr(err, w)
		return
	}

	err = hs.templates.ExecuteTemplate(w, "article-card.html", a)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
