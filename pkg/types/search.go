package types

type QueryBody struct {
	Term string `json:"term"`
}

type SearchQuery struct {
	Source     []string  `json:":source"`
	From       int       `json:"from"`
	MaxResults int16     `json:"max_results"`
	SearchType string    `json:"search_type"`
	Query      QueryBody `json:"query"`
}

type SearchByTermBody struct {
	Term string `json:"term"`
}
