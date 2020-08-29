package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jaeger-tracing-go-service/config"
	"github.com/jaeger-tracing-go-service/routes"
)

func main()  {
	// Database
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}