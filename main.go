package main

import (
	"net/http"

	"github.com/disto-stack/j-mails-indexer/pkg/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main()  {
	r := chi.NewRouter()

	r.Use(middleware.Logger)


	r.Route("/api/v1", func(r chi.Router) {
		routes.Indexer(r)
	})

	http.ListenAndServe(":3000", r)
}

func hello(w http.ResponseWriter, r*http.Request) {
	w.Write([]byte("Hello world"))
}