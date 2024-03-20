package types

type SearchResult struct {
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		Hits []struct {
			ID     string `json:"_id"`
			Source Place  `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type Place struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
}

type Foodcorts struct {
	Name     string `json:"name"`
	Total    string `json:"total"`
	Places   string `json:"places"`
	PrevPage string `json:"prev_page"`
	NextPage string `json:"next_page"`
	LastPage string `json:"last_page"`
}
