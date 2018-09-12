package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maxkrasnov/ru.maxkrasnov.api/pkg/api/models"
	"net/http"
	"strconv"
)

func (cr *Controller) Projects(c *gin.Context)  {
	var project models.Project
	var err error
	page := c.Query("page")
	p := 1
	if page != "" {
		p, err = strconv.Atoi(page)
	}
	if err == nil {
		projects := project.GetAll(cr.DB, p)
		c.JSON(http.StatusOK, projects)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": false})
	}
}