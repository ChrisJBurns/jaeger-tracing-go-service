package main

import (
	"github.com/gin-gonic/gin"
	"jaeger-tracing-go-service/config"
	"jaeger-tracing-go-service/routes"
	"log"
)

func main() {
	// Set client options
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
