package htmx

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/not4sure/news-service-test-task/internal/app"
)

//go:embed views/*
var views embed.FS

type HTMXServer struct {
	app       app.Application
	templates *template.Template
}

func NewHTMXServer(application app.Application) HTMXServer {
	// TODO: handle error
	templates := template.Must(template.ParseFS(views, "views/*"))

	return HTMXServer{
		app:       application,
		templates: templates,
	}
}

func (hs *HTMXServer) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", hs.HandleMainPage)

	router.HandleFunc("GET /article/{id}/editor", hs.ArticleEditorHandler)

	router.HandleFunc("GET /article/create-form", hs.ArticleCreatorHandler)

	router.HandleFunc("POST /article", hs.ArticleCreateHandler)

	router.HandleFunc("PUT /article/{id}", hs.ArticleUpdateHandler)

	router.HandleFunc("GET /article/{id}", hs.ArticleGetHandler)

	router.HandleFunc("DELETE /article/{id}", hs.ArticleDeleteHandler)
}
