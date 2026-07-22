package shared

import (
	"main/internal/domain"

	"github.com/gin-gonic/gin"
)

func GetJWTPayload(c *gin.Context) (*domain.JWTPayload, error) {
	value, exist := c.Get("user")
	if !exist {
		return nil, ErrUnauthorized
	}

	user, ok := value.(*domain.JWTPayload)

	if !ok {
		return nil, ErrUnauthorized
	}

	return user, nil
}
