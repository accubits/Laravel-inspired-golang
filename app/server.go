package main

import (
	"myproject/app/config"
	"myproject/app/middleware"
	"myproject/app/routes"
	"fmt"
	"log"
)

func main() {

	config, err := config.New()
	if err != nil {
		log.Panicln("Configuration error", err)
	}
	fmt.Println(config)
	router := routes.New()
	router = middleware.New(router, config)
	router = routes.API(router, config)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
