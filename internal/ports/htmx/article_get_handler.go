package htmx

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (hs *HTMXServer) ArticleGetHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	a, err := hs.app.GetByID(r.Context(), uid)
	err = hs.templates.ExecuteTemplate(w, "article-card.html", a)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
