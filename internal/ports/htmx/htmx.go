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
	tmpl := template.New("").Funcs(template.FuncMap{
		"dict": func(values ...any) map[string]any {
			dict := make(map[string]any)
			for i := 0; i < len(values); i += 2 {
				key := values[i].(string)
				dict[key] = values[i+1]
			}
			return dict
		},
	})

	templates := template.Must(tmpl.ParseFS(views, "views/*"))

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
