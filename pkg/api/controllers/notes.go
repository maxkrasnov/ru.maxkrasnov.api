package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maxkrasnov/ru.maxkrasnov.api/pkg/api/models"
	"net/http"
	"strconv"
)

func (cr *Controller) Notes(c *gin.Context)  {
	var note models.Note
	var err error
	page := c.Query("page")
	p := 1
	if page != "" {
		p, err = strconv.Atoi(page)
	}
	if err == nil {
		notes := note.GetAll(cr.DB, p)
		c.JSON(http.StatusOK, notes)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": false})
	}
}

func (cr *Controller) Note(c *gin.Context)  {
	var note models.Note
	note.GetByCode(c.Param("code"), cr.DB)
	if note.ID > 0 {
		c.JSON(http.StatusOK, note)
	} else {
		c.JSON(http.StatusNotFound, note)
	}
}