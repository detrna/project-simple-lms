package shared

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, logger Logger, err error) {
	logger.Warn(err.Error())
	switch {
	case errors.Is(err, ErrEmailTaken):
		c.JSON(http.StatusConflict, err)

	case errors.Is(err, ErrCredentialsIncorrect):
		c.JSON(http.StatusConflict, err)

	case errors.Is(err, ErrBadRequest):
		c.JSON(http.StatusBadRequest, err)

	case errors.Is(err, ErrRecordNotFound):
		c.JSON(http.StatusBadRequest, err)

	default:
		c.JSON(http.StatusInternalServerError, err)
	}
}

var (
	ErrEmailTaken           = errors.New("email was already taken")
	ErrCredentialsIncorrect = errors.New("incorrect email or password")
	ErrBadRequest           = errors.New("bad request")
	ErrRecordNotFound       = errors.New("couldn't find any record of requested data")
)
