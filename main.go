package main

import (
	"net/http"
	"os"

	"github.com/disto-stack/j-mails-indexer/pkg/handlers"
	"github.com/disto-stack/j-mails-indexer/pkg/routes"
	"github.com/disto-stack/j-mails-indexer/pkg/services"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

var (
	configService     *services.ConfigService
	zincSearchService *services.ZincsearchService
	indexerHandler    *handlers.IndexerHandler
	emailHandler      *handlers.EmailHandler
	emailRoutes       *routes.EmailRoutes
)

type envVariables struct {
	zincsearchUrl string
}

func main() {
	setupDependencies()

	args := os.Args
	if len(args) > 1 {
		dir := args[1]
		indexerHandler.IndexFromDir(dir)
	}

	setupServer()
}

func setupDependencies() {
	configService = &services.ConfigService{}
	configService.SetUrlsFromEnv()

	// Services
	zincSearchService = &services.ZincsearchService{}
	zincSearchService.SetDependencies(configService)

	// Handlers
	indexerHandler = &handlers.IndexerHandler{}
	indexerHandler.SetDependencies(configService, zincSearchService)

	emailHandler = &handlers.EmailHandler{}
	emailHandler.SetDependencies(configService, zincSearchService)

	// Routes
	emailRoutes = &routes.EmailRoutes{}
	emailRoutes.SetDependencies(emailHandler)
}

func setupServer() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Route("/api/v1/", func(r chi.Router) {
		emailRoutes.MountEmailRouter(r)
	})

	http.ListenAndServe(":3000", r)
}
