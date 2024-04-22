package api

import (
	"aitrainer-api/config"
	"aitrainer-api/docs"
	"aitrainer-api/internal/api/books"
	"aitrainer-api/internal/auth"
	"aitrainer-api/internal/middleware"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/time/rate"
)

func InitRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	if gin.Mode() == gin.ReleaseMode {
		r.Use(middleware.Security())
		r.Use(middleware.Xss())
	}
	r.Use(middleware.Cors())
	r.Use(middleware.RateLimiter(rate.Every(1*time.Minute), 60)) // 60 requests per minute

	authHandlers := auth.RegisterAuthHandler(cfg)
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", books.Healthcheck)
		v1.GET("/books", middleware.APIKeyAuth(cfg), books.FindBooks)
		v1.POST("/books", middleware.APIKeyAuth(cfg), middleware.JWTAuth(), books.CreateBook)
		v1.GET("/books/:id", middleware.APIKeyAuth(cfg), books.FindBook)
		v1.PUT("/books/:id", middleware.APIKeyAuth(cfg), books.UpdateBook)
		v1.DELETE("/books/:id", middleware.APIKeyAuth(cfg), books.DeleteBook)

		v1.POST("/login", middleware.APIKeyAuth(cfg), authHandlers.LoginHandler)
		v1.POST("/register", middleware.APIKeyAuth(cfg), authHandlers.RegisterHandler)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
