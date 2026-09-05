package middleware

import (
	"main/internal/pkg"
	"main/internal/shared/pagination"

	"github.com/gin-gonic/gin"
)

func HandlePagination(defaultLimit int, maxLimit int, logger pkg.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging pagination.Pagination

		if err := ctx.ShouldBindQuery(&paging); err != nil {
			paging.Limit = defaultLimit
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = defaultLimit
		}

		if paging.Limit >= maxLimit {
			paging.Limit = maxLimit
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		paging.Offset = (paging.Page - 1) * paging.Limit

		ctx.Set("pagination", paging)
		ctx.Next()
	}
}
