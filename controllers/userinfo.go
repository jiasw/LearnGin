package controllers

import (
	"gorm.io/gorm"
	"net/http"
	"visiontest/dtos"
	"visiontest/models"

	"github.com/gin-gonic/gin"
)

// UserInfoController is a controller for user info
type UserInfoController struct {
	Router *gin.Engine
}

// GetUserInfo is a handler function for getting user info
func (uic *UserInfoController) GetUserInfo(c *gin.Context) {

	db := c.MustGet("DB").(*gorm.DB)
	var users []models.UserInfo
	if err := db.Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	dtos.SuccessResponse(c, users)
}

func (uic *UserInfoController) SaveUserInfo(c *gin.Context) {
	userinfo := models.UserInfo{}
	if err := c.ShouldBindJSON(&userinfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("DB").(*gorm.DB)
	if err := db.Create(&userinfo).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
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
