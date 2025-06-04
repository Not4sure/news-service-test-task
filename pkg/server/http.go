package server

import "net/http"

type RegisterFn func(router *http.ServeMux)

func RunServer(register RegisterFn) {
	RunServerOnAddr(":8080", register)
}

func RunServerOnAddr(addr string, register RegisterFn) {
	router := http.NewServeMux()
	register(router)

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	server.ListenAndServe()
}
