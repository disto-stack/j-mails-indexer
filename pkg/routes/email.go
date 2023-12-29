package routes

import (
	"github.com/disto-stack/j-mails-indexer/pkg/handlers"
	"github.com/go-chi/chi"
)

type EmailRoutes struct {
	emailHandler *handlers.EmailHandler
}

func (e *EmailRoutes) SetDependencies(em *handlers.EmailHandler) {
	e.emailHandler = em
}

func (e *EmailRoutes) MountEmailRouter(r chi.Router) {
	r.Mount("/email", e.mailRoutes())
}

func (e *EmailRoutes) mailRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/search", e.emailHandler.SearchByTerm)

	return r
}
