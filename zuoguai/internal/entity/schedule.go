package entity

import "gorm.io/gorm"

type Schedule struct {
	ID         int    `gorm:"primaryKey,column:id'" json:"id"`
	UserId     int    `gorm:"user_id" json:"user_id"`
	Title      string `gorm:"title" json:"title"`
	Content    string `gorm:"content" json:"content"`
	Status     int    `gorm:"status" json:"status"`
	StartTime  int64  `gorm:"start_time" json:"start_time"`
	EndTime    int64  `gorm:"end_time" json:"end_time"`
	CreateTime int64  `gorm:"create_time" json:"create_time"`
	UpdateTime int64  `gorm:"update_time" json:"update_time"`
	Mark       string `gorm:"mark" json:"mark"`
}

func (s *Schedule) TableName() string {
	return "t_schedule"
}

type ScheduleEntity interface {
	CreateSchedule(db *gorm.DB, schedule *Schedule) (err error)
	UpdateSchedule(db *gorm.DB, schedule *Schedule) (err error)
	DeleteSchedule(db *gorm.DB, schedule *Schedule) (err error)
	GetScheduleById(db *gorm.DB, id int) (schedule Schedule, err error)
	GetScheduleListByCond(db *gorm.DB, options ...ScheduleFindOptionFn) (schedules []*Schedule, err error)
	GetScheduleCountByCond(db *gorm.DB, options ...ScheduleFindOptionFn) (count int64, err error)
}

type IScheduleEntity struct {
}

func NewScheduleEntity() ScheduleEntity {
	return &IScheduleEntity{}
}

func (s *IScheduleEntity) CreateSchedule(db *gorm.DB, schedule *Schedule) (err error) {
	err = db.Model(&Schedule{}).Create(schedule).Error
	return err
}

func (s *IScheduleEntity) UpdateSchedule(db *gorm.DB, schedule *Schedule) (err error) {
	err = db.Model(&Schedule{}).Where("id =?", schedule.ID).Updates(schedule).Error
	return err
}
func (s *IScheduleEntity) DeleteSchedule(db *gorm.DB, schedule *Schedule) (err error) {
	err = db.Model(&Schedule{}).Where("id =?", schedule.ID).Delete(&schedule).Error
	return err
}
func (s *IScheduleEntity) GetScheduleById(db *gorm.DB, id int) (schedule Schedule, err error) {
	err = db.Model(&Schedule{}).Where("id =?", id).First(&schedule).Error
	return
}

type ScheduleFindOptionFn func(db *gorm.DB)
type ScheduleFindOptionInstance struct {
}

var ScheduleFindOptionIns = &ScheduleFindOptionInstance{}

func (s *IScheduleEntity) GetScheduleListByCond(db *gorm.DB, options ...ScheduleFindOptionFn) (schedules []*Schedule, err error) {
	db = db.Model(&Schedule{})
	for _, option := range options {
		option(db)
	}
	err = db.Find(&schedules).Error
	return
}

func (s *IScheduleEntity) GetScheduleCountByCond(db *gorm.DB, options ...ScheduleFindOptionFn) (count int64, err error) {
	db = db.Model(&Schedule{})
	for _, option := range options {
		option(db)
	}
	err = db.Count(&count).Error
	return
}
func (s *ScheduleFindOptionInstance) WithId(id int) ScheduleFindOptionFn {
	return func(db *gorm.DB) {
		db = db.Where("id =?", id)
	}
}

func (s *ScheduleFindOptionInstance) WithUserId(userId int) ScheduleFindOptionFn {
	return func(db *gorm.DB) {
		db = db.Where("user_id =?", userId)
	}
}

func (s *ScheduleFindOptionInstance) WithTitle(title string) ScheduleFindOptionFn {
	return func(db *gorm.DB) {
		db = db.Where("title =?", title)
	}
}

func (s *ScheduleFindOptionInstance) WithContent(content string) ScheduleFindOptionFn {
	return func(db *gorm.DB) {
		db = db.Where("content =?", content)
	}
}
func (s *ScheduleFindOptionInstance) WithStatus(status int) ScheduleFindOptionFn {
	return func(db *gorm.DB) {
		db = db.Where("status =?", status)
	}
}

func (s *ScheduleFindOptionInstance) WithTime(startTime int64, endTime int64) ScheduleFindOptionFn {
	return func(db *gorm.DB) {
		db = db.Where("start_time >= ? && end_time <= ?", startTime, endTime)
	}
}
func (s *ScheduleFindOptionInstance) WithCreateTime(startTime int64) ScheduleFindOptionFn {
	return func(db *gorm.DB) {
		db = db.Where("create_time >=?", startTime)
	}
}

func (s *ScheduleFindOptionInstance) WithMark(mark string) ScheduleFindOptionFn {
	return func(db *gorm.DB) {
		db = db.Where("mark =?", mark)
	}
}

func (s *ScheduleFindOptionInstance) WithSkipPage(pageStart, pageEnd int) ScheduleFindOptionFn {
	return func(db *gorm.DB) {
		db = db.Offset(pageStart).Limit(pageEnd)
	}
}
