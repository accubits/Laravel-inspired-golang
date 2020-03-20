package controller

import (
	// "conferencing/app/model"
	"myproject/app/config"
	"myproject/app/services"

	"github.com/gin-gonic/gin"
)

func HomePage(config *config.Config) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		c.String(200, services.Login(config))
	}
	
	return gin.HandlerFunc(fn)
}
