package main

import (
	"net/http"

	"github.com/not4sure/news-service-test-task/internal/app"
	ports "github.com/not4sure/news-service-test-task/internal/ports/htmx"
	"github.com/not4sure/news-service-test-task/pkg/server"
)

func main() {
	application := app.NewApplication()

	server.RunServer(func(router *http.ServeMux) {
		hs := ports.NewHTMXServer(application)
		hs.RegisterRoutes(router)
	})
}
