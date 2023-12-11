package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Indexer(r chi.Router) {
	r.Mount("/indexer", indexRoutes())
}

func indexRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func (w http.ResponseWriter, r *http.Request)  {
			w.Write([]byte("Hello from indexer"))
	})

	return r
}