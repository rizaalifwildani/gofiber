package models

// Pagination represents pagination information.
type Pagination struct {
	CurrentPage int         `json:"currentPage"`
	TotalData   int         `json:"totalData"`
	PerPage     int         `json:"perPage"`
	NextPage    int         `json:"nextPage"`
	Data        interface{} `json:"data"`
}
