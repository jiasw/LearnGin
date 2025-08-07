package controllers

import (
	"visiontest/dtos"

	"github.com/gin-gonic/gin"
)

// Home 首页
// @Summary 首页
// @Description 这是首页的描述
// @Produce json
// @Tags 首页
// @Success 200 {string} string "ok"
// @Router / [get]
func Home(c *gin.Context) {
	msg := "Hello World"
	dtos.SuccessResponse(c, msg)
}
