package service

import (
	"sync"
	"zuoguai/internal/entity"
	"zuoguai/internal/service/metadata"
	"zuoguai/internal/service/schedule"
	"zuoguai/internal/service/user"
)

var services *Services
var once sync.Once

func GetServices() *Services {
	once.Do(func() {
		services = NewServices()
	})
	return services
}

type Services struct {
	UserService     metadata.UserServiceInterface
	ScheduleService metadata.ScheduleServiceInterface
}

func NewServices() *Services {
	return &Services{
		UserService:     user.NewUserService(entity.NewUserEntity()),
		ScheduleService: schedule.NewScheduleService(entity.NewScheduleEntity(), entity.NewUserEntity()),
	}
}
