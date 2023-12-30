package main

import (
	"os"

	"github.com/disto-stack/j-mails-indexer/pkg/handlers"
	"github.com/disto-stack/j-mails-indexer/pkg/services"
)

var (
	indexerHandler handlers.IndexerHandler
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
}

func setupDependencies() {
	config := &services.Config{}
	config.SetUrlsFromEnv()

	// Services
	zincSearchService := services.ZincsearchService{}
	zincSearchService.SetDependencies(config)

	// Handlers
	indexerHandler = handlers.IndexerHandler{}
	indexerHandler.SetDependencies(config, &zincSearchService)
}
