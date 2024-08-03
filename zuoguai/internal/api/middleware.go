package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"zuoguai/internal/db"
	"zuoguai/internal/entity"
	"zuoguai/pkg/constant"
	"zuoguai/utils"
)

func AuthMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		//验证Token存不存在，token正不正确
		jwtToken := c.GetHeader("Authorization") //
		jwtToken, ok := strings.CutPrefix(jwtToken, "Bearer ")
		// var myClaims *utils.MyClaims
		if ok {
			if claims, err := utils.ParseJwtToken(jwtToken, string(utils.DefaultSecretKey)); err == nil {
				userID, err := strconv.Atoi(claims.UserId)
				if err != nil || !checkUserId(userID) {
					c.JSON(http.StatusUnauthorized, "用户被封禁或不存在")
					c.Abort()
					return
				}
				c.Set(constant.USERID, userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, "token过期")
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, "未携带token")
		c.Abort()
	}
}

func checkUserId(userId int) bool {
	if user, err := entity.GetEntities().UserEntity.GetUserByUserId(db.GetDB(), userId); err != nil {
		return false
	} else if user.ID == 0 {
		return false
	} else if user.Status == constant.UserStatusBlock {
		return false
	}

	return true
}
