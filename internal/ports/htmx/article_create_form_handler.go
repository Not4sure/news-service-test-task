package htmx

import (
	"net/http"
)

func (hs *HTMXServer) ArticleCreateFormHandler(w http.ResponseWriter, r *http.Request) {
	err := hs.templates.ExecuteTemplate(w, "article-create-form.html", nil)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
