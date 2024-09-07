package types

type Pagination struct {
	Page  int
	Limit int
}

type GetParam struct {
	Query      map[string]interface{}
	Pagination Pagination
}

type PaginatedResponse struct {
	Results []interface{} `json:"result"`
	Page    int           `json:"page"`
	Limit   int           `json:"limit"`
	Total   int           `json:"total"`
}
