package main

import (
	"log"

	"github.com/gin-gonic/gin"
	config "github.com/syadav214/config"
	routes "github.com/syadav214/routes"
)

func main() {
	config.Connect()
	// Init Router
	router := gin.Default()
	// Route Handlers / Endpoints
	routes.Routes(router)
	log.Fatal(router.Run(":4747"))
}
