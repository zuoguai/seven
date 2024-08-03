package entity

import "sync"

type Entities struct {
	UserEntity     UserEntity
	ScheduleEntity ScheduleEntity
}

var (
	entities Entities
	once     sync.Once
)

func GetEntities() Entities {
	once.Do(newEntities)
	return entities
}

func newEntities() {
	entities = Entities{
		UserEntity:     NewUserEntity(),
		ScheduleEntity: NewScheduleEntity(),
	}
}
