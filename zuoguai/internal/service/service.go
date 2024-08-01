package service

import (
	"zuoguai/internal/entity"
	"zuoguai/internal/service/metadata"
	"zuoguai/internal/service/user"
)

type Services struct {
	UserService metadata.UserServiceInterface
}

func NewServices() *Services {
	return &Services{
		UserService: user.NewUserService(entity.NewUserEntity()),
	}
}
