package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r gin.IRouter) {

	api := r.Group("/api")
	user := api.Group("/user")
	{
		user.POST("/login", Login)
		user.POST("/register", Register)
		user.POST("/update", AuthMiddleWare(), UpdateUser)

	}

	schedule := api.Group("/schedule", AuthMiddleWare())
	{
		schedule.POST("", AddSchedule)
		schedule.GET("", GetSchedule)
	}
}
