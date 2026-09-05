package controller

import (
	"main/internal/http/response"
	"main/internal/http/validator"
	"main/internal/modules/class/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl ClassController) Update(c *gin.Context) {
	var input dto.UpdateClassRequest

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		validator.HandleValidationError(c, ctrl.logger, err)
		return
	}

	if err := c.ShouldBindUri(&input); err != nil {
		validator.HandleValidationError(c, ctrl.logger, err)
		return
	}

	ctx := c.Request.Context()
	result, err := ctrl.uc.Update(ctx, &input)

	if err != nil {
		response.Error(c, ctrl.logger, err)
		return
	}

	dto := response.ResponseDTO[dto.ClassResponse]{
		StatusCode: http.StatusOK,
		Data:       *dto.DomainToResponse(*result),
	}

	response.Success(c, dto)
}
