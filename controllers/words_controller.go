package controllers

import (
	"api_golang_ia/services"
	"github.com/gin-gonic/gin"
)

type WordsController struct {}

func (pc WordsController) Index(c *gin.Context) {
	service := services.IaServices{}

	c.JSON(200, service.GetWords())
}