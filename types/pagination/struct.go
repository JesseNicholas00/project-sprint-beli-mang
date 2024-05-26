package pagination

type Page struct {
	Limit  *int   `json:"limit"`
	Offset *int   `json:"offset"`
	Total  *int64 `json:"total"`
}
