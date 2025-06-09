package main

import (
	"context"
	"net/http"

	ports "github.com/not4sure/news-service-test-task/internal/ports/htmx"
	"github.com/not4sure/news-service-test-task/internal/service"
	"github.com/not4sure/news-service-test-task/pkg/server"
)

func main() {
	ctx := context.Background()

	application := service.NewApplication(ctx)

	server.RunServer(func(router *http.ServeMux) {
		hs := ports.NewHTMXServer(application)
		hs.RegisterRoutes(router)
	})
}
