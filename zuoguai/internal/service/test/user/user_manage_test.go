package user_test

import (
	"fmt"
	"testing"
	"zuoguai/internal/config"
	"zuoguai/internal/service"
	"zuoguai/internal/service/metadata"
)

func TestGetUserList(t *testing.T) {

	config.GetConfigs("../../../.././app.yaml")
	service := service.NewServices()
	resp, err := service.UserService.GetUserList(metadata.GetUserInfoListRequest{
		Username: "",
		// Role:     "角色10",
		// Phone:     "电话11",
		// Email:     "邮件97",
		// Status:    2,
		PageStart: 3,
		PageSize:  2,
	})
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(resp.List); i++ {
		fmt.Println(resp.List[i])
	}
	fmt.Println(resp.Total)

}
