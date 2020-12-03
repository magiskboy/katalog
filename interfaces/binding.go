package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/magiskboy/katalog/core"
	"github.com/magiskboy/katalog/interfaces/rest"
	"net/http"
)

// GetHTTPBinding get a http application
func GetHTTPBinding() *gin.Engine {

	router := gin.New()

	router.Use(gin.Recovery())
	if env := core.Getenv("ENV", "development"); env == "development" {
		router.Use(gin.Logger())
	}

	router.GET("/health", func(context *gin.Context) {
		context.Writer.WriteHeader(http.StatusNoContent)
	})

	category := router.Group("/categories")
	{
		category.POST("", rest.CreateSellerCategoryHandler)
	}

	return router
}
