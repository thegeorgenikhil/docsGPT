package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thegeorgenikhil/docsGPT/controllers"
)

func EmbeddingRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/api/query", controllers.QueryFromEmbeddedDocument())
}
