package htmx

import (
	"net/http"

	"github.com/google/uuid"
)

func (hs *HTMXServer) ArticleEditFormHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	a, err := hs.app.GetByID(r.Context(), uid)
	if err != nil {
		hs.handleAppErr(err, w)
		return
	}

	err = hs.templates.ExecuteTemplate(w, "article-edit-form.html", a)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
