package htmx

import (
	"html/template"
	"net/http"

	"github.com/not4sure/news-service-test-task/internal/app"
	"github.com/not4sure/news-service-test-task/internal/domain/article"
)

// HTMXServer serves htmx view for application
type HTMXServer struct {
	app       app.Application
	templates *template.Template
}

// NewHTMXServer creates HTMXServer for given Application implementation.
func NewHTMXServer(application app.Application) HTMXServer {
	templates, err := readTemplates()
	if err != nil {
		panic(err)
	}

	return HTMXServer{
		app:       application,
		templates: templates,
	}
}

func (hs *HTMXServer) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", hs.HandleMainPage)

	router.HandleFunc("GET /article/{id}/edit", hs.ArticleEditFormHandler)

	router.HandleFunc("GET /article/create", hs.ArticleCreateFormHandler)

	router.HandleFunc("POST /article", hs.ArticleCreateHandler)

	router.HandleFunc("PUT /article/{id}", hs.ArticleUpdateHandler)

	router.HandleFunc("GET /article/{id}", hs.ArticleGetHandler)

	router.HandleFunc("DELETE /article/{id}", hs.ArticleDeleteHandler)
}

func (hs *HTMXServer) handleAppErr(err error, w http.ResponseWriter) {
	if err == nil {
		return
	}

	if err == article.ErrArticleNotFound {
		http.Error(w, "article not found", http.StatusNotFound)
		return
	}

	http.Error(w, "failed to list articles", http.StatusInternalServerError)
}
