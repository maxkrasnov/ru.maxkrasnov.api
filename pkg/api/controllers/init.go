package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	DB *gorm.DB
}

func (cr *Controller) Index(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Привет! Есть вопросы, пиши на me@maxkrasnov.ru",
		"version": "1.0.0",
	})
}