package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maxkrasnov/ru.maxkrasnov.api/pkg/api/models"
	"net/http"
)

func (cr *Controller) Feedback(c *gin.Context)  {
	feedback := models.Feedback{}
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}
	cr.DB.Create(&feedback)
	if feedback.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": true})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": false})
	}
}
