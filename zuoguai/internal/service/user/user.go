package user

import "zuoguai/internal/entity"

type UserService struct {
	userEntity entity.UserEntity
}

func NewUserService(user entity.UserEntity) *UserService {
	return &UserService{
		userEntity: user,
	}
}
