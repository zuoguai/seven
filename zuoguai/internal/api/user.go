package api

import (
	"github.com/gin-gonic/gin"
	"zuoguai/internal/pkg"
	"zuoguai/internal/service"
	"zuoguai/internal/service/metadata"
)

func Login(c *gin.Context) {
	var (
		req      metadata.LoginRequest
		resp     metadata.LoginResponse
		response pkg.HttpResponse
		err      error
	)

	if err = c.BindJSON(&req); err != nil {
		response.ResponseWithError(c, pkg.CodeLoginErr, "登录参数绑定出错")
		return
	}

	if resp, err = service.GetServices().UserService.LoginUser(req); err == nil {
		response.ResponseWithData(c, resp)
	} else {
		response.ResponseWithError(c, pkg.CodeLoginErr, err.Error())
	}

}

func Register(c *gin.Context) {
	var (
		req      metadata.RegisterRequest
		resp     metadata.RegisterResponse
		response pkg.HttpResponse
		err      error
	)

	if err = c.BindJSON(&req); err != nil {
		response.ResponseWithError(c, pkg.CodeLoginErr, "注册参数绑定出错")
		return
	}

	if resp, err = service.GetServices().UserService.RegisterUser(req); err == nil {
		response.ResponseWithData(c, resp)
	} else {
		response.ResponseWithError(c, pkg.CodeLoginErr, err.Error())
	}

}

func UpdateUser(c *gin.Context) {
	//var (
	//	req      metadata.UpdateUserInfoRequest
	//	response pkg.HttpResponse
	//	err      error
	//)
	//
	//if err = c.BindJSON(&req); err != nil {
	//	response.ResponseWithError(c, pkg.CodeLoginErr, "登录参数绑定出错")
	//}
	//
	//if data, err := service.GetServices().UserService.(req); err == nil {
	//	response.ResponseWithData(c, data)
	//} else {
	//	response.ResponseWithError(c, pkg.CodeLoginErr, "登录服务出错")
	//}

}
