package controllers

import (
	"net/http"
	"visiontest/dtos"
	"visiontest/infrastructure/databasehelper"
	"visiontest/infrastructure/repositories"
	"visiontest/models"

	"github.com/gin-gonic/gin"
)

// UserInfoController is a controller for user info
type UserInfoController struct {
	Router *gin.Engine
}

// GetUserInfo is a handler function for getting user info
func (uic *UserInfoController) GetUserInfo(c *gin.Context) {
	userRep := repositories.NewUserInfoRepository(databasehelper.GetInstance().DB)
	users, total, _ := userRep.Paginate(1, 10)
	dtos.SuccessResponse(c, gin.H{
		"data":  users,
		"total": total,
		"page":  1,
		"limit": 10,
	})
}

func (uic *UserInfoController) SaveUserInfo(c *gin.Context) {
	userinfo := models.UserInfo{}
	if err := c.ShouldBindJSON(&userinfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := databasehelper.GetInstance().DB
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
