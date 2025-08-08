package dtos

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse struct {
	Code    int         `json:"code"`           // 使用 HTTP 状态码而非自定义业务码
	Message string      `json:"message"`        // 人类可读的提示信息
	Data    interface{} `json:"data"`           // 实际业务数据
	Meta    interface{} `json:"meta,omitempty"` // 分页等元数据
}

// 常用状态响应构造器
func SuccessResponseWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

// 最基础的成果响应构造器
func SuccessResponse(c *gin.Context) {
	c.JSON(http.StatusOK, ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    nil,
	})
}

func ErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusOK, ApiResponse{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	})
}
