package pagination

import "math"

type PaginationResponse struct {
	Page       int  `json:"page"`
	Limit      int  `json:"limit"`
	HasNext    bool `json:"hasNext"`
	HasPrev    bool `json:"hasPrev"`
	TotalItems int  `json:"totalItems"`
	TotalPages int  `json:"totalPages"`
}

type PaginationInput struct {
	Page   int
	Offset int
	Limit  int
}

type PaginationRequest struct {
	Page  int
	Limit int
}

type Items interface {
	[]any
	int
}

func GetPaginationResponse(pagination PaginationRequest, totalItems int) *PaginationResponse {
	return &PaginationResponse{
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		HasNext:    (pagination.Page * pagination.Limit) <= totalItems,
		HasPrev:    pagination.Page != 0,
		TotalItems: totalItems,
		TotalPages: int(math.Ceil(float64(totalItems) / float64(pagination.Page))),
	}
}
