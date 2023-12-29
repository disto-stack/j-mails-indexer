package types

type SearchHits struct {
	Total struct {
		Value int `json:"value"`
	} `json:"total"`
	Hits []struct {
		Source Email `json:"_source"`
	} `json:"hits"`
}

type ZincsearchApiResponse struct {
	Took int        `json:"took"`
	Hits SearchHits `json:"hits"`
}
