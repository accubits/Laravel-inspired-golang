package main

import (
	"log"
	"myproject/app/config"
	"myproject/app/database/migration"
	"myproject/app/middleware"
	"myproject/app/routes"
)



func main() {
	config, err := config.New()
	if err != nil {
		log.Panicln("Configuration error", err)
	}
	migration.New(config)
	router := routes.New()
	router = middleware.New(router, config)
	router = routes.API(router, config)
	router.Run(":" + config.Env.GetString("server.port"))
}
