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

type Pagination struct {
	Offset int `form:"-"`
	Page   int `form:"page" binding:"required"`
	Limit  int `form:"limit" binding:"required"`
}

func GetPaginationResponse(pagination Pagination, totalItems int) *PaginationResponse {
	return &PaginationResponse{
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		HasNext:    (pagination.Page * pagination.Limit) <= totalItems,
		HasPrev:    pagination.Page != 0,
		TotalItems: totalItems,
		TotalPages: int(math.Ceil(float64(totalItems) / float64(pagination.Page))),
	}
}
