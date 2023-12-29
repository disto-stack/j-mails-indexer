package handlers_test

import (
	"testing"

	"github.com/disto-stack/j-mails-indexer/pkg/handlers"
	"github.com/disto-stack/j-mails-indexer/pkg/services"
)

func TestReadFromDir(t *testing.T) {
	configService := services.ConfigService{}
	configService.SetUrlsFromEnv()

	zincsearchService := services.ZincsearchService{}
	zincsearchService.SetDependencies(&configService)

	indexer := handlers.IndexerHandler{}
	indexer.SetDependencies(&configService, &zincsearchService)

	indexer.IndexFromDir("/home/disto/new_bulk_data")
}
