package entity

import "sync"

type Entities struct {
	UserEntity UserEntity
}

var (
	entities Entities
	once     sync.Once
)

func GetEntities() Entities {
	once.Do(newEntityes)
	return entities
}

func newEntityes() {
	entities = Entities{
		UserEntity: NewUserEntity(),
	}
}
