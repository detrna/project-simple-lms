package shared

import (
	"errors"
	"main/internal/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Error string `json:"error"`
}

func HandleError(c *gin.Context, logger pkg.Logger, err error) {
	logger.Warn(err.Error())

	error := ResponseError{
		Error: err.Error(),
	}

	switch {
	case errors.Is(err, ErrEmailTaken) || errors.Is(err, ErrSystemIDTaken):
		c.JSON(http.StatusConflict, error)
		return

	case errors.Is(err, ErrCredentialsIncorrect):
		c.JSON(http.StatusConflict, error)
		return

	case errors.Is(err, ErrBadRequest):
		c.JSON(http.StatusBadRequest, error)
		return

	case errors.Is(err, ErrRecordNotFound):
		c.JSON(http.StatusBadRequest, error)
		return

	case errors.Is(err, ErrIncorrectOTP):
		c.JSON(http.StatusBadRequest, error)
		return

	case errors.Is(err, ErrForbidden):
		c.AbortWithStatusJSON(http.StatusForbidden, error)
		return

	case errors.Is(err, ErrUnauthorized):
		c.AbortWithStatusJSON(http.StatusUnauthorized, error)
		return

	default:
		c.JSON(http.StatusInternalServerError, error)
		return
	}
}

var (
	ErrEmailTaken           = errors.New("email was already taken")
	ErrSystemIDTaken        = errors.New("systemID was already taken")
	ErrCredentialsIncorrect = errors.New("incorrect email or password")
	ErrBadRequest           = errors.New("bad request")
	ErrRecordNotFound       = errors.New("couldn't find any record of requested data")
	ErrIncorrectOTP         = errors.New("incorrect otp code")
	ErrUnauthorized         = errors.New("request unauthorized")
	ErrForbidden            = errors.New("request forbidden")
)
