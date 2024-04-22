package main

import (
	"aitrainer-api/config"
	"aitrainer-api/internal/api"
	"aitrainer-api/internal/cache"
	"aitrainer-api/internal/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8001
// @BasePath  /api/v1

// @securityDefinitions.apikey JwtAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	cfg := config.LoadConfig()

	cache.InitRedis(cfg)
	database.ConnectDatabase(cfg)

	if cfg.IsDev() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := api.InitRouter(cfg)

	if err := r.Run(fmt.Sprintf(":%v", cfg.Port)); err != nil {
		log.Fatal(err)
	}
}
