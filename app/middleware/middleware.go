package middleware

import (
	"myproject/app/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Init initalize the middlewares
func New(router *gin.Engine, config *config.Config) *gin.Engine {
	router.Use(cors.Default())
	return router
}
