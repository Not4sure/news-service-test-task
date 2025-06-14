package htmx

import (
	"net/http"

	"github.com/google/uuid"
)

func (hs *HTMXServer) ArticleDeleteHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = hs.app.DeleteArticle(r.Context(), uid)
	if err != nil {
		hs.handleAppErr(err, w)
		return
	}

	// Proper response would be 204 No Content, but htmx won't handle it properly
	w.WriteHeader(http.StatusOK)
}
