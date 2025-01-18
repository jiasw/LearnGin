package controllers

import (
	"fmt"
	"net/http"
	"visiontest/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserInfoController is a controller for user info
type UserInfoController struct {
	Router *gin.Engine
	Db     *gorm.DB
}

var Usercontroller UserInfoController

// GetUserInfo is a handler function for getting user info
func (uic *UserInfoController) GetUserInfo(c *gin.Context) {
	firstuser := models.UserInfo{}
	uic.Db.First(&firstuser)
	fmt.Println("GetUserInfo", firstuser)
	c.JSON(http.StatusOK, gin.H{
		"message": "User info deleted successfully",
		"data":    firstuser,
	})
}

func (uic *UserInfoController) SaveUserInfo(c *gin.Context) {

	userinfo := models.UserInfo{
		Name: "test",

		Email: "test@test.com",
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User info saved successfully",
		"data":    userinfo,
	})

}

func (uic *UserInfoController) DeleteUserInfo(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "User info deleted successfully",
	})

}
