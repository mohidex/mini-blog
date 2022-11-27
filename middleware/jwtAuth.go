package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/mini-blog/helper"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := helper.ValidateJWT(ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authentication required",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
