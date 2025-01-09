package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// UserInfoController is a controller for user info
type UserInfoController struct {
}

var Usercontroller UserInfoController

var userlist []UserInfo = []UserInfo{
	{Name: "John Doe", Age: 30},
	{Name: "Jane Smith", Age: 25},
}

// GetUserInfo is a handler function for getting user info
func (uic *UserInfoController) GetUserInfo(c *gin.Context) {
	req, _ := c.Get("request")
	fmt.Println("request:", req)

	c.JSON(http.StatusOK, userlist)
}

func (uic *UserInfoController) SaveUserInfo(c *gin.Context) {
	// TODO: save user info to database
	userlist = append(userlist, UserInfo{Name: "jiasw", Age: 20})

	c.JSON(http.StatusOK, gin.H{
		"message": "User info saved successfully",
	})
}
