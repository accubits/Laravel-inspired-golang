package controller

import (
	"myproject/app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Register(config *config.Config) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		
		c.JSON(http.StatusOK, gin.H{
			"msg":"registered",
		})

	}
	return gin.HandlerFunc(fn)

}
func Login(config *config.Config) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		
		c.JSON(http.StatusOK, gin.H{
			"msg":"login",
		})

	}

	return gin.HandlerFunc(fn)
}
