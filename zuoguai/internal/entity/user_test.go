package entity_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"zuoguai/internal/config"
	"zuoguai/internal/db"
	"zuoguai/internal/entity"
)

func TestCreateUser(t *testing.T) {
	config.GetConfigs("../.././app.yaml")
	for i := 0; i < 1000; i++ {
		user := entity.User{
			Username:  fmt.Sprintf("作怪%d", i),
			Password:  fmt.Sprintf("密码%d", rand.Int31()%100),
			Email:     fmt.Sprintf("邮件%d", rand.Int31()%100),
			Phone:     fmt.Sprintf("电话%d", rand.Int31()%100),
			Nickname:  fmt.Sprintf("昵称%d", rand.Int31()%100),
			Role:      fmt.Sprintf("角色%d", rand.Int31()%100),
			Sex:       i%2 + 1,
			Status:    i % 3,
			IsVisitor: 1,
			Avatar:    fmt.Sprintf("作怪%d", rand.Int31()%100),
			Address:   fmt.Sprintf("作怪%d", rand.Int31()%100),
			Birthday:  time.Now().Unix(),
			Remark:    fmt.Sprintf("个性签名%d", rand.Int31()%100),
		}
		entity.NewUserEntity().CreateUser(db.GetDB(), &user)
		fmt.Println(user)
	}
}
