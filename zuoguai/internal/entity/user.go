package entity

import "gorm.io/gorm"

type User struct {
	ID        int    `gorm:"column:id,primary_key"`
	Username  string `gorm:"username"`
	Password  string `gorm:"password"`
	Role      string `gorm:"role"`
	Token     string `gorm:"token"`
	Email     string `gorm:"email"`
	Phone     string `gorm:"phone"`
	Avatar    string `gorm:"avatar"`
	Nickname  string `gorm:"nickname"`
	Sex       int    `gorm:"sex"`
	Age       int    `gorm:"age"`
	Birthday  int    `gorm:"birthday"`
	Address   string `gorm:"address"`
	Remark    string `gorm:"remark"`
	Status    int    `gorm:"status"`
	CreateAt  int64  `gorm:"create_at"`
	UpdateAt  int64  `gorm:"update_at"`
	DeleteAt  int64  `gorm:"delete_at"`
	IsVisitor int    `gorm:"is_visitor"`
}

func (u *User) TableName() string {
	return "user"
}

type UserEntity interface {
	CreateUser(db *gorm.DB, user *User) error
	UpdateUser(db *gorm.DB, user *User) error
	DeleteUser(db *gorm.DB, user *User) error
	FindUserListByCond(db *gorm.DB, options ...UserFindOptionFn) ([]User, error)
}

type iUserEntity struct {
}

func NewUserEntity() UserEntity {
	return &iUserEntity{}
}

func (i *iUserEntity) CreateUser(db *gorm.DB, user *User) (err error) {
	err = db.Model(&User{}).Create(user).Error
	return
}

func (i *iUserEntity) UpdateUser(db *gorm.DB, user *User) (err error) {
	err = db.Model(&User{}).Updates(&user).Error
	return err
}

func (i *iUserEntity) DeleteUser(db *gorm.DB, user *User) (err error) {
	err = db.Model(&User{}).Delete(user).Error
	return err
}

func (i *iUserEntity) FindUserListByCond(db *gorm.DB, options ...UserFindOptionFn) (list []User, err error) {
	db = db.Model(&User{})
	for _, option := range options {
		option(db)
	}

	err = db.Scan(&list).Error

	return
}

type UserFindOptionFn func(db *gorm.DB)

type UserFindOptionInstance struct {
}

var UserFindOptionIns = &UserFindOptionInstance{}

func (_ UserFindOptionInstance) WithID(id int) UserFindOptionFn {
	return func(db *gorm.DB) {
		db.Where("id = ?", id)
	}
}
func (_ UserFindOptionInstance) WithUsername(username string) UserFindOptionFn {
	return func(db *gorm.DB) {
		db.Where("username = ?", username)
	}
}

// func (_ UserFindOptionInstance) WithPassword(password string) UserFindOptionFn {
// 	return func(db *gorm.DB) {
// 		db.Where("password = ?", password)
// 	}
// }

func (_ UserFindOptionInstance) WithRole(role string) UserFindOptionFn {
	return func(db *gorm.DB) {
		db.Where("role = ?", role)
	}
}

// func (_ UserFindOptionInstance) WithToken(token string) UserFindOptionFn {
// 	return func(db *gorm.DB) {
// 		db.Where("token = ?", token)
// 	}
// }

func (_ UserFindOptionInstance) WithAge(age int) UserFindOptionFn {
	return func(db *gorm.DB) {
		db.Where("avatar = ?", age)
	}
}

func (_ UserFindOptionInstance) WithEmail(email string) UserFindOptionFn {
	return func(db *gorm.DB) {
		db.Where("email = ?", email)
	}
}

func (_ UserFindOptionInstance) WithPhone(phone string) UserFindOptionFn {
	return func(db *gorm.DB) {
		db.Where("phone = ?", phone)
	}
}
func (_ UserFindOptionInstance) WithNickname(nickname string) UserFindOptionFn {
	return func(db *gorm.DB) {
		db.Where("nickname = ?", nickname)
	}
}
func (_ UserFindOptionInstance) WithSex(sex int) UserFindOptionFn {
	return func(db *gorm.DB) {
		db.Where("sex = ?", sex)
	}
}
func (_ UserFindOptionInstance) WithStatus(status int) UserFindOptionFn {
	return func(db *gorm.DB) {
		db.Where("status = ?", status)
	}
}
func (_ UserFindOptionInstance) WithBirthdayPeriod(startTime, endTime int64) UserFindOptionFn {
	return func(db *gorm.DB) {
		db.Where("birthday >= ? AND birthday <= ?", startTime, endTime)
	}
}
func (_ UserFindOptionInstance) WithCreateAtPeriod(startTime, endTime int64) UserFindOptionFn {
	return func(db *gorm.DB) {
		db.Where("create_at >= ? AND create_at <= ?", startTime, endTime)
	}
}
func (_ UserFindOptionInstance) WithUpdateAtPeriod(startTime, endTime int64) UserFindOptionFn {
	return func(db *gorm.DB) {
		db.Where("update_at >= ? AND update_at <= ?", startTime, endTime)
	}
}
func (_ UserFindOptionInstance) WithDeleteAtPeriod(startTime, endTime int64) UserFindOptionFn {
	return func(db *gorm.DB) {
		db.Where("delete_at >= ? AND delete_at <= ?", startTime, endTime)
	}
}
func (_ UserFindOptionInstance) WithIsVisitor(isVisitor int) UserFindOptionFn {
	return func(db *gorm.DB) {
		db.Where("is_visitor = ?", isVisitor)
	}
}
