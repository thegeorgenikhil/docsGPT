package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thegeorgenikhil/docsGPT/controllers"
)

func DocumentRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/api/documents", controllers.GetDocuments())
	incomingRoutes.POST("/api/documents/upload", controllers.UploadDocumentToEmbed())
}
