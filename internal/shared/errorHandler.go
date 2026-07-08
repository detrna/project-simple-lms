package shared

import (
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrEmailTaken):
		log.Printf("Error: %v\nStack:\n%s", err, debug.Stack())
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})

	case errors.Is(err, ErrCredentialsIncorrect):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		log.Printf("Error: %v\nStack:\n%s", err, debug.Stack())

	case errors.Is(err, ErrBadRequest):
		c.JSON(http.StatusBadRequest, err)
		log.Printf("Error: %v\nStack:\n%s", err, debug.Stack())

	case errors.Is(err, ErrRecordNotFound):
		c.JSON(http.StatusBadRequest, err)
		log.Printf("Error: %v\nStack:\n%s", err, debug.Stack())

	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		log.Printf("Error: %v\nStack:\n%s", err, debug.Stack())
	}
}

var (
	ErrEmailTaken           = errors.New("email was already taken")
	ErrCredentialsIncorrect = errors.New("incorrect email or password")
	ErrBadRequest           = errors.New("bad request")
	ErrRecordNotFound       = errors.New("couldn't find any record of requested data")
)
