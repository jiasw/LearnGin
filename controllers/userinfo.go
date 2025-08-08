package controllers

import (
	"net/http"
	"strconv"
	"visiontest/dtos"
	"visiontest/infrastructure/databasehelper"
	"visiontest/infrastructure/repositories"
	"visiontest/models"

	"github.com/gin-gonic/gin"
)

type UserInfoController struct {
	Router *gin.Engine
}

// @Summary 获取人员分页列表
// @Description 获取人员分页列表
// @Tags 用户信息
// @Produce json
// @Param page query int true "页码"
// @Param limit query int true "每页数量"
// @Security BearerAuth
// @Success 200 {object} dtos.ApiResponse
// @Router /userlist [get]
func (uic *UserInfoController) GetUserInfoList(c *gin.Context) {
	userRep := repositories.NewUserInfoRepository(databasehelper.GetInstance().DB)
	pagestr, limitstr := c.Query("page"), c.Query("limit")
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	limit, err := strconv.Atoi(limitstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	users, total, _ := userRep.Paginate(page, limit)
	dtos.SuccessResponseWithData(c, gin.H{
		"userlist": users,
		"total":    total,
	})
}

// @Summary 获取人员详情
// @Description 获取人员详情
// @Tags 用户信息
// @Produce json
// @Param id query int true "用户ID"
// @Success 200 {object} dtos.ApiResponse
// @Router /userinfo [get]
func (uic *UserInfoController) GetUserInfoByID(c *gin.Context) {
	idstr := c.Query("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		dtos.ErrorResponse(c, err.Error())
		return
	}
	userRep := repositories.NewUserInfoRepository(databasehelper.GetInstance().DB)
	userinfo, err := userRep.GetByID(uint(id))
	if err != nil {
		dtos.ErrorResponse(c, err.Error())
		return
	}
	dtos.SuccessResponseWithData(c, userinfo)
}

// CreateUserInfo @Summary 创建人员
// @Description 创建人员
// @Tags 用户信息
// @Produce json
// @Param user body models.UserInfo true "用户信息"
// @Success 200 {object} dtos.ApiResponse
// @Router /createUser [post]
func (uic *UserInfoController) CreateUserInfo(c *gin.Context) {
	userinfo := models.UserInfo{}
	if err := c.ShouldBindJSON(&userinfo); err != nil {
		dtos.ErrorResponse(c, err.Error())
		return
	}
	db := databasehelper.GetInstance().DB
	if err := db.Create(&userinfo).Error; err != nil {
		dtos.ErrorResponse(c, err.Error())
		return
	}
	dtos.SuccessResponse(c)
}

// UpdateUserInfo @Summary 修改人员信息
// @Description 修改人员信息
// @Tags 用户信息
// @Produce json
// @Param user body models.UserInfo true "用户信息"
// @Success 200 {object} dtos.ApiResponse
// @Router /updateUser [post]
func (uic *UserInfoController) UpdateUserInfo(c *gin.Context) {
	userinfo := models.UserInfo{}
	if err := c.ShouldBindJSON(&userinfo); err != nil {
		dtos.ErrorResponse(c, err.Error())
		return
	}
	db := databasehelper.GetInstance().DB
	if err := db.Save(&userinfo).Error; err != nil {
		dtos.ErrorResponse(c, err.Error())
		return
	}
	dtos.SuccessResponse(c)
}

// DeleteUserByID @Summary 删除人员
// @Description 删除人员
// @Tags 用户信息
// @Produce json
// @Param id formData int true "用户ID"
// @Success 200 {object} dtos.ApiResponse
// @Router /delUser [post]
func (uic *UserInfoController) DeleteUserByID(c *gin.Context) {
	idstr := c.PostForm("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		dtos.ErrorResponse(c, err.Error())
		return
	}
	db := databasehelper.GetInstance().DB
	if err := db.Where("id = ?", id).Delete(&models.UserInfo{}).Error; err != nil {
		dtos.ErrorResponse(c, err.Error())
		return
	}
	dtos.SuccessResponse(c)
}
