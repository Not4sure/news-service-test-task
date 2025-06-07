package htmx

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (hs *HTMXServer) ArticleEditorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("editing article")

	uid, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	a, err := hs.app.GetByID(r.Context(), uid)
	err = hs.templates.ExecuteTemplate(w, "article-editor.html", a)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
