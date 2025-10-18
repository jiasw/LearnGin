package controllers

import (
	"visiontest/dtos"
	"visiontest/infrastructure/databasehelper"
	jwthelper "visiontest/infrastructure/jwtHelper"
	"visiontest/infrastructure/repositories"

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
	msg := "欢迎你来到首页"
	dtos.SuccessResponseWithData(c, msg)
}

// Login 登录
// @Summary 登录
// @Description 这是登录的描述
// @Produce json
// @Tags 首页
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "ok"
// @Router /login [post]
func Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	userRep := repositories.NewUserInfoRepository(databasehelper.GetInstance().DB)
	user, err := userRep.Where("name = ?", username).First()
	if err != nil {
		dtos.ErrorResponse(c, "密码错误")
	}
	if user == nil {
		dtos.ErrorResponse(c, "密码错误")
	}
	if user.Password == password {
		token, err := jwthelper.GenerateToken(uint(user.ID),
			user.Name,
			user.Address,
			"1")
		if err != nil {
			dtos.ErrorResponse(c, "登录失败")
		}
		dtos.SuccessResponseWithData(c, token)
	} else {
		dtos.ErrorResponse(c, "登录失败")
	}
}
