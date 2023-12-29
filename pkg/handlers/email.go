package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/disto-stack/j-mails-indexer/pkg/services"
	"github.com/disto-stack/j-mails-indexer/pkg/types"
)

type EmailHandler struct {
	configService     *services.ConfigService
	zincSearchService *services.ZincsearchService
}

func (e *EmailHandler) SetDependencies(c *services.ConfigService, z *services.ZincsearchService) {
	e.configService = c
	e.zincSearchService = z
}

func (e *EmailHandler) SearchByTerm(w http.ResponseWriter, r *http.Request) {
	searchTerm := types.SearchByTermBody{}
	err := json.NewDecoder(r.Body).Decode(&searchTerm)

	if err != nil {
		log.Fatalln(err)
	}

	searchQuery := types.SearchQuery{
		Source:     []string{},
		From:       0,
		MaxResults: 20,
		SearchType: "match",
		Query: types.QueryBody{
			Term: searchTerm.Term,
		},
	}

	zincResponse := e.zincSearchService.SearchByTerm(searchQuery)
	w.WriteHeader(int(zincResponse.Code))
	json.NewEncoder(w).Encode(zincResponse)
}
