package controllers

import (
	"go-backend/configs"
	"go-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {
	var person models.Person

	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := configs.DB.Create(&person)

	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Person created successfully",
		"data":    person,
	})
}
