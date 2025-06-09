package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

type RegisterFn func(router *http.ServeMux)

func RunServer(register RegisterFn) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	RunServerOnAddr(addr, register)
}

func RunServerOnAddr(addr string, register RegisterFn) {
	l := zerolog.New(os.Stdout)

	router := http.NewServeMux()
	register(router)

	server := &http.Server{
		Addr:    addr,
		Handler: NewLoggerMiddleware(l, router),
	}

	server.ListenAndServe()
}

func LoggingMiddleware(l zerolog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			l.Debug().Msg("hello")
			fmt.Println("lol")

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
