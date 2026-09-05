package controller

import (
	"main/internal/http/response"
	"main/internal/http/validator"
	classdto "main/internal/modules/class/dto"
	shareddto "main/internal/shared/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl ClassController) GetByID(c *gin.Context) {
	params := validator.ValidateParams[shareddto.IDParams](c, ctrl.logger)
	if params == nil {
		return
	}

	ctx := c.Request.Context()
	class, err := ctrl.uc.GetByID(ctx, params.ID)

	if err != nil {
		response.Error(c, ctrl.logger, err)
		return
	}

	dto := response.ResponseDTO[classdto.ClassResponse]{
		StatusCode: http.StatusOK,
		Data:       *classdto.DomainToResponse(*class),
	}

	response.Success(c, dto)
}
