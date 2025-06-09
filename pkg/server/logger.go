package server

import (
	"net/http"

	"github.com/rs/zerolog"
)

type LoggerMiddleware struct {
	l       zerolog.Logger
	handler http.Handler
}

func (lm *LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lm.l.Info().
		Str("method", r.Method).
		Str("uri", r.RequestURI).
		Msg("http request")

	lm.handler.ServeHTTP(w, r)
}

func NewLoggerMiddleware(l zerolog.Logger, hndl http.Handler) *LoggerMiddleware {
	return &LoggerMiddleware{
		l:       l,
		handler: hndl,
	}
}
