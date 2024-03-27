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
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Phone    string   `json:"phone"`
	Location Location `json:"location"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Foodcorts struct {
	Name     string  `json:"name"`
	Total    int     `json:"total"`
	Places   []Place `json:"places"`
	PrevPage int     `json:"prev_page"`
	NextPage int     `json:"next_page"`
	LastPage int     `json:"last_page"`
}

type PageData struct {
	Name        string
	Total       int
	Prev        int
	Next        int
	PerPage     int
	CurrentPage int
	List        []Place
}
