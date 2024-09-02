package http_api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/kerim-dauren/smart-terminal/api/http_api/v1"
	"github.com/kerim-dauren/smart-terminal/internal/manager"
	"github.com/kerim-dauren/smart-terminal/pkg/loggerx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewHttpRouter(logger loggerx.Logger, managers *manager.Managers) *gin.Engine {
	handler := gin.New()

	if gin.Mode() == gin.DebugMode {
		handler.Use(gin.Logger())
	}
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	handler.GET("/health", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	// Routers
	apiV1 := handler.Group("/api/v1")
	{
		v1.Register(apiV1, managers)
	}

	return handler
}
