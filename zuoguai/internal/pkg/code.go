package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess             = 0     // http请求成功
	CodeParamErr            = 10002 // 请求参数不合法
	CodeRegisterErr         = 10003 // 注册错误
	CodeLoginErr            = 10003 // 登录错误
	CodeLogoutErr           = 10004 // 登出错误
	CodeGetUserInfoErr      = 10005 // 获取用户信息错误
	CodeUpdateUserInfoErr   = 10006 // 更新用户信息错误
	CodeAddScheduleError    = 89809
	CodeGetScheduleError    = 898009
	CodeUpdateScheduleError = 839809
)

// HttpResponse http独立请求返回结构体
type HttpResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseWithError http请求返回处理函数
func (resp *HttpResponse) ResponseWithError(c *gin.Context, code int, msg string) {
	resp.Code = code
	resp.Msg = msg
	c.JSON(http.StatusInternalServerError, resp)
}

func (resp *HttpResponse) ResponseSuccess(c *gin.Context) {
	resp.Code = CodeSuccess
	resp.Msg = "success"
	c.JSON(http.StatusOK, resp)
}

func (resp *HttpResponse) ResponseWithData(c *gin.Context, data interface{}) {
	resp.Code = CodeSuccess
	resp.Msg = "success"
	resp.Data = data
	c.JSON(http.StatusOK, resp)
}
