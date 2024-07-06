package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func PersonRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/persons", controllers.CreatePerson)
}
