package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	//成功Code
	SUCCESS int = 0
	FAIL int = 1
)

//请求成功返回内容
func Success(context *gin.Context, v interface{}) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"Code":SUCCESS,
		"Message":"success",
		"data":v,
	})
}

//请求失败返回内容
func Fail(context *gin.Context, v interface{}) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"Code":FAIL,
		"Message":"fail",
		"data":v,
	})
}
