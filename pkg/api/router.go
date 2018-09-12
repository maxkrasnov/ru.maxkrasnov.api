package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maxkrasnov/ru.maxkrasnov.api/pkg/api/controllers"
	"github.com/maxkrasnov/ru.maxkrasnov.api/pkg/api/models"
)
//
func InitAPI()  {
	router := gin.Default()

	db := models.InitDB()
	clr := controllers.Controller{
		DB: db,
	}

	router.Use(CORSMiddleware())

	router.GET("/", clr.Index)
	router.GET("/notes", clr.Notes)
	router.GET("/notes/:code", clr.Note)
	router.GET("/resume", clr.Resume)
	router.GET("/projects", clr.Projects)
	router.POST("/feedback", clr.Feedback)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "NOT_FOUND", "message": "Route not found"})
	})
	router.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}