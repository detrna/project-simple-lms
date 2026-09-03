package shared

import (
	"errors"
	"main/internal/pkg"
	apperrors "main/internal/shared/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func HandleErrorV2(c *gin.Context, logger pkg.Logger, appErr apperrors.AppError, err error) {
	logger.WarnSkip(1, err.Error())

	c.JSON(appErr.StatusCode, gin.H{
		"error":   err.Error(),
		"message": appErr.Message,
	})
}

func HandleError(c *gin.Context, logger pkg.Logger, err error) {
	logger.WarnSkip(1, err.Error())

	error := ResponseError{
		Error: err.Error(),
	}

	switch {
	case errors.Is(err, ErrEmailTaken) || errors.Is(err, ErrSystemIDTaken):
		c.JSON(http.StatusConflict, error)
		return

	case errors.Is(err, ErrCredentialsIncorrect):
		c.JSON(http.StatusUnauthorized, error)
		return

	case errors.Is(err, ErrBadRequest):
		c.JSON(http.StatusBadRequest, error)
		return

	case errors.Is(err, ErrRecordNotFound):
		c.JSON(http.StatusNotFound, error)
		return

	case errors.Is(err, ErrRedisRecordNotFound):
		c.JSON(http.StatusNotFound, error)
		return

	case errors.Is(err, ErrIncorrectOTP):
		c.JSON(http.StatusUnauthorized, error)
		return

	case errors.Is(err, ErrForbidden):
		c.AbortWithStatusJSON(http.StatusForbidden, error)
		return

	case errors.Is(err, ErrUnauthorized):
		c.AbortWithStatusJSON(http.StatusUnauthorized, error)
		return

	case errors.Is(err, ErrTokenMissing):
		c.AbortWithStatusJSON(http.StatusUnauthorized, error)
		return

	case errors.Is(err, ErrInvalidToken):
		c.AbortWithStatusJSON(http.StatusUnauthorized, error)
		return

	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, error)
		return
	}
}

var (
	ErrEmailTaken           = errors.New("email was already taken")
	ErrSystemIDTaken        = errors.New("systemID was already taken")
	ErrCredentialsIncorrect = errors.New("incorrect email or password")
	ErrBadRequest           = errors.New("bad request")
	ErrRecordNotFound       = errors.New("couldn't find any record of requested data")
	ErrRedisRecordNotFound  = errors.New("couldn't find any record of requested redis data")
	ErrIncorrectOTP         = errors.New("incorrect otp code")
	ErrUnauthorized         = errors.New("request unauthorized")
	ErrForbidden            = errors.New("request forbidden")
	ErrTokenMissing         = errors.New("authorization header did not exist")
	ErrInvalidToken         = errors.New("invalid token")
)
