package controllers

import (
	"api_golang_ia/models"
	"github.com/gin-gonic/gin"
)

type HomeController struct {}

func (hc HomeController) Index(c *gin.Context) {
	c.JSON(200, models.Message{
		Message: "Aqui é a mensagem de retorno do home controller!",
	})
}