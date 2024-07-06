package main

import (
	"go-backend/configs"
	"go-backend/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.DatabaseInit()
}

func main() {
	r := gin.Default()

	routes.PersonRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()

}
