package controller

import (
	"fmt"
	"main/internal/http/response"
	"main/internal/http/validator"
	"main/internal/shared/dto"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl ClassController) Delete(c *gin.Context) {
	var params dto.IDParams
	if err := c.ShouldBindUri(&params); err != nil {
		validator.HandleValidationError(c, ctrl.logger, err)
		return
	}

	fmt.Print("TESTING")
	fmt.Print(params.ID)

	ctx := c.Request.Context()
	err := ctrl.uc.Delete(ctx, params.ID)

	if err != nil {
		response.Error(c, ctrl.logger, err)
		return
	}

	c.Status(http.StatusNoContent)
}
