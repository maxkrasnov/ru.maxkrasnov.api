package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maxkrasnov/ru.maxkrasnov.api/pkg/api/models"
	"net/http"
)

func (cr *Controller) Resume(c *gin.Context)  {
	var resume models.Resume
	resume.Get(cr.DB)
	c.JSON(http.StatusOK, resume)
}
