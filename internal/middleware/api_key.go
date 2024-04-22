package middleware

import (
	"aitrainer-api/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIKeyAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == cfg.ApiKey {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
		}
	}
}
