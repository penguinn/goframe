package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess = 0
	CodeFailure = 1
)

type Render struct {
	Result  interface{} `json:"result,omitempty"`  // 成功返回的结构体
	Message string      `json:"message,omitempty"` // 错误返回信息
	Status  int         `json:"status"`            // 返回code，0表示成功
	Success bool        `json:"success"`           // 成功返回true，失败返回false
}

// JSON渲染错误，最终会渲染为{"success":false, "error": "错误信息"}
func renderError(c *gin.Context, err error) {
	data := Render{
		Success: false,
		Status:  CodeFailure,
		Message: err.Error(),
	}

	c.JSON(200, data)
}

// JSON渲染信息，最终会渲染为{"success":true, "result": 渲染信息}
func renderInfo(c *gin.Context, info interface{}) {
	data := Render{
		Success: true,
		Result:  info,
		Status:  CodeSuccess,
	}

	c.JSON(200, data)
}

// JSON渲染没有额外信息的成功响应，最终会渲染为{"success":true}
func renderSuccess(c *gin.Context) {
	renderInfo(c, nil)
}

func renderInvalidRequest(c *gin.Context) {
	renderError(c, errors.New("请求参数验证错误！"))
}

func renderInternalServerError(c *gin.Context) {
	renderError(c, errors.New("服务器内部发生错误！"))
}
