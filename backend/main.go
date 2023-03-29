package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thegeorgenikhil/docsGPT/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}

	r := gin.Default()

	r.Use(cors.Default())

	routes.DocumentRoutes(r)
	routes.EmbeddingRoutes(r)

	log.Fatal(r.Run(":" + port))
}
