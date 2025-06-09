package htmx

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/not4sure/news-service-test-task/internal/domain/article"
)

func (hs *HTMXServer) ArticleDeleteHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = hs.app.Delete(r.Context(), uid)
	if err != nil {
		if err == article.ErrArticleNotFound {
			http.Error(w, "article not found", http.StatusNotFound)
			return
		}

		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	// Proper response would be 204 No Content, but htmx won't handle it properly
	w.WriteHeader(http.StatusOK)
}
