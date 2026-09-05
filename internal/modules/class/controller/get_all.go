package controller

import (
	"main/internal/http/response"
	"main/internal/modules/class/dto"
	"main/internal/shared/pagination"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl ClassController) GetAll(c *gin.Context) {
	paging := c.MustGet("pagination").(pagination.Pagination)

	ctx := c.Request.Context()
	classes, total, err := ctrl.uc.GetAll(ctx, &paging)

	if err != nil {
		response.Error(c, ctrl.logger, err)
		return
	}

	dto := response.ResponseDTO[[]dto.ClassResponse]{
		StatusCode: http.StatusOK,
		Data:       *dto.DomainToResponseBatch(*classes),
		Pagination: pagination.GetPaginationResponse(paging, total),
	}

	response.Success(c, dto)
}
