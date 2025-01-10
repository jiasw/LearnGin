package controllers

import (
	"visiontest/dtos"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	msg := "Hello World"
	dtos.SuccessResponse(c, msg)
}
