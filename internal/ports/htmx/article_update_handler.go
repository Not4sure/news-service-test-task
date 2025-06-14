package htmx

import (
	"net/http"

	"github.com/google/uuid"
)

func (hs *HTMXServer) ArticleUpdateHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	title := r.Form.Get("title")
	content := r.Form.Get("content")

	a, err := hs.app.UpdateArticle(r.Context(), uid, title, content)
	if err != nil {
		hs.handleAppErr(err, w)
		return
	}

	err = hs.templates.ExecuteTemplate(w, "article-card.html", a)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
