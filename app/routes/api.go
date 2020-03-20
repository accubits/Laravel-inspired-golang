package routes

import (
	c "myproject/app/controller"
	"github.com/gin-gonic/gin"
	"myproject/app/config"
)


func New() *gin.Engine{
	r := gin.Default()
	return r
}

//API handles all the routes
func API(r *gin.Engine,config *config.Config) *gin.Engine {

	r.GET("/ping", c.HomePage(config))
	
	return r
}

