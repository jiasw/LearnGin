package dtos

import (
	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	response := ApiResponse{
		Code:    200,
		Message: "success",
		Data:    data,
	}
	c.JSON(200, response)
}

func ErrorResponse(c *gin.Context, code int, message string) {
	response := ApiResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	}
	c.JSON(code, response)
}
