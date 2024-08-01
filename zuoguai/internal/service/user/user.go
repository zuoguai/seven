package user

import (
	"zuoguai/internal/entity"
	"zuoguai/internal/service/metadata"
)

type UserService struct {
	userEntity entity.UserEntity
}

func NewUserService(user entity.UserEntity) metadata.UserServiceInterface {
	return UserService{
		userEntity: user,
	}
}
