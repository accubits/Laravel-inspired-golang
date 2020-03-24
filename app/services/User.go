package services

import (
	"myproject/app/config"

	"github.com/gin-gonic/gin"
)

func Login(config *config.Config, c *gin.Context) string {

	return config.Env.GetString("server.port")
}
