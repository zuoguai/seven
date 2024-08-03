package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"zuoguai/internal/pkg"
	"zuoguai/internal/service"
	"zuoguai/internal/service/metadata"
	"zuoguai/pkg/constant"
)

func AddSchedule(c *gin.Context) {
	var (
		req      metadata.ScheduleRequest
		resp     metadata.ScheduleResponse
		response pkg.HttpResponse
		err      error
	)

	if err = c.BindJSON(&req); err != nil {
		fmt.Println(req, " err", err)

		response.ResponseWithError(c, pkg.CodeAddScheduleError, "添加任务失败，参数绑定出错")
		return
	}
	ctx := context.WithValue(context.Background(), constant.USERID, c.GetInt(constant.USERID))

	if resp, err = service.GetServices().ScheduleService.AddASchedule(ctx, req); err == nil {
		response.ResponseWithData(c, resp)
	} else {
		response.ResponseWithError(c, pkg.CodeAddScheduleError, err.Error())
	}

}

func GetSchedule(c *gin.Context) {
	var (
		req      metadata.ScheduleRequest
		resp     metadata.ScheduleResponse
		response pkg.HttpResponse
		err      error
	)

	if err = c.BindJSON(&req); err != nil {
		response.ResponseWithError(c, pkg.CodeGetScheduleError, "获取任务失败，参数绑定出错")
		return
	}
	ctx := context.WithValue(context.Background(), constant.USERID, c.GetInt(constant.USERID))

	if resp, err = service.GetServices().ScheduleService.GetASchedule(ctx, req); err == nil {
		response.ResponseWithData(c, resp)
	} else {
		response.ResponseWithError(c, pkg.CodeGetScheduleError, err.Error())
	}

}
