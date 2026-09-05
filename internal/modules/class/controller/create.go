package controller

import (
	"main/internal/http/response"
	"main/internal/http/validator"
	"main/internal/modules/class/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl ClassController) Create(c *gin.Context) {
	var body dto.CreateClassRequest
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		validator.HandleValidationError(c, ctrl.logger, err)
		return
	}

	ctx := c.Request.Context()
	result, err := ctrl.uc.Create(ctx, &body)

	if err != nil {
		response.Error(c, ctrl.logger, err)
		return
	}

	dto := response.ResponseDTO[dto.ClassResponse]{
		StatusCode: http.StatusCreated,
		Data:       *dto.DomainToResponse(*result),
	}

	response.Success(c, dto)
}
