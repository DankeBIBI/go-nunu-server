package api

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	NotFound = func(msg string) string { return "数据未找到：" + msg }
)

func Success(ctx *gin.Context, data interface{}, msg ...string) {
	Msg := "操作成功"
	if data == nil {
		data = map[string]interface{}{}
	}
	if len(msg) != 0 {
		Msg = msg[0]
	}
	res := Response{
		Code: 1,
		Msg:  Msg,
		Data: data,
	}
	ctx.JSON(200, res)
}
func Fail(ctx *gin.Context, data interface{}, msg ...string) {
	Msg := "操作失败"
	if data == nil {
		data = map[string]interface{}{}
	}
	if len(msg) != 0 {
		Msg = msg[0]
	}
	res := Response{
		Code: 0,
		Msg:  Msg,
		Data: data,
	}
	ctx.JSON(500, res)
}
