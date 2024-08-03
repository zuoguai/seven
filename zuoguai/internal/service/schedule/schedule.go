package schedule

import (
	"errors"
	"golang.org/x/net/context"
	"time"
	"zuoguai/internal/db"
	"zuoguai/internal/entity"
	"zuoguai/internal/service/metadata"
	"zuoguai/pkg/constant"
)

type ScheduleService struct {
	entity.ScheduleEntity
	entity.UserEntity
}

func NewScheduleService(schedule entity.ScheduleEntity, user entity.UserEntity) metadata.ScheduleServiceInterface {
	return ScheduleService{
		schedule,
		user,
	}
}
func (scheduleService ScheduleService) GetASchedule(ctx context.Context, req metadata.ScheduleRequest) (resp metadata.ScheduleResponse, err error) {
	if req.UserId == 0 {
		err = errors.New("用户不存在")
	}
	var schedule entity.Schedule
	ownerId := ctx.Value(constant.USERID).(int)
	owner, err := scheduleService.GetUserByUserId(db.GetDB(), ownerId)
	if err != nil {
		err = errors.New("用户不存在")
		return
	}

	if owner.ID != req.UserId {
		if owner.Role != constant.UserRoleAdmin {
			err = errors.New("不能查看别人的日程")
			return
		}
	}

	schedule, err = scheduleService.GetScheduleById(db.GetDB(), req.ID)
	if err != nil || schedule.Status == constant.ScheduleStatusDelete || schedule.Status == constant.ScheduleStatusBlock {
		err = errors.New("日程不存在")
	}
	if owner.ID != schedule.UserId {
		if owner.Role != constant.UserRoleAdmin {
			err = errors.New("不能查看别人的日程")
			return
		}
	}
	resp.Schedule = schedule
	return
}

func (scheduleService ScheduleService) AddASchedule(ctx context.Context, req metadata.ScheduleRequest) (resp metadata.ScheduleResponse, err error) {
	if req.UserId == 0 {
		err = errors.New("用户不存在")
	}
	var schedule entity.Schedule
	ownerId := ctx.Value(constant.USERID).(int)
	owner, err := scheduleService.GetUserByUserId(db.GetDB(), ownerId)
	if err != nil {
		err = errors.New("用户不存在")
		return
	}

	if owner.ID != req.UserId {
		if owner.Role != constant.UserRoleAdmin {
			err = errors.New("不能修改别人的日程")
			return
		}
	}

	req.ID = 0
	if len(req.Title) < 2 {
		err = errors.New("标题过短")
		return
	}
	if len(req.Title) > 50 {
		err = errors.New("标题过长")
		return
	}
	if len(req.Content) < 2 {
		err = errors.New("内容过短")
		return
	}
	if len(req.Content) > 200 {
		err = errors.New("内容过长")
		return
	}
	if req.StartTime == 0 || req.EndTime < req.StartTime {
		err = errors.New("开始时间不能为空 且结束时间必须大于开始时间")
		return
	}
	schedule = req.Schedule
	schedule.UserId = req.UserId
	schedule.CreateTime = time.Now().Unix()
	schedule.UpdateTime = time.Now().Unix()
	schedule.Status = constant.ScheduleStatusNormal
	schedule.ID = 0
	err = scheduleService.CreateSchedule(db.GetDB(), &schedule)
	resp.Schedule = schedule
	if err != nil {
		err = errors.New("创建日程失败")
	}
	return
}

func (scheduleService ScheduleService) UpdateASchedule(ctx context.Context, req metadata.ScheduleRequest) (resp metadata.ScheduleResponse, err error) {
	if req.UserId == 0 {
		err = errors.New("用户不存在")
	}
	var schedule entity.Schedule
	ownerId := ctx.Value(constant.USERID).(int)
	owner, err := scheduleService.GetUserByUserId(db.GetDB(), ownerId)
	if err != nil {
		err = errors.New("用户不存在")
		return
	}

	if owner.ID != req.UserId {
		if owner.Role != constant.UserRoleAdmin {
			err = errors.New("不能修改别人的日程")
			return
		}
	}

	if req.ID == 0 {
		err = errors.New("日程不存在")
		return
	}

	schedule, err = scheduleService.GetScheduleById(db.GetDB(), req.ID)
	if err != nil || schedule.Status == constant.ScheduleStatusDelete || schedule.Status == constant.ScheduleStatusBlock {
		if owner.Role != constant.UserRoleAdmin {
			err = errors.New("日程不存在")
			return
		}
	}

	if schedule.UserId != owner.ID {
		if owner.Role != constant.UserRoleAdmin {
			err = errors.New("不能修改别人的日程")
			return
		}
	}

	if len(req.Title) < 2 && len(req.Title) > 50 {
		schedule.Title = req.Title

	}
	if len(req.Content) < 2 && len(req.Content) > 200 {
		schedule.Content = req.Content
	}

	if req.StartTime > 0 && req.StartTime < req.EndTime {
		schedule.StartTime = req.StartTime
		schedule.EndTime = req.EndTime
	}
	if req.Status != 0 {
		switch req.Status {
		case constant.ScheduleStatusNormal:
			schedule.Status = req.Status
		case constant.ScheduleStatusBlock:
			schedule.Status = req.Status
		case constant.ScheduleStatusDelete:
			schedule.Status = req.Status
		case constant.ScheduleStatusSuccess:
			schedule.Status = req.Status
		case constant.ScheduleStatusFail:
			schedule.Status = req.Status
		case constant.ScheduleStatusExpire:
			schedule.Status = req.Status
		default:
			schedule.Status = constant.ScheduleStatusNormal
		}
	}
	if req.Mark != "" {
		schedule.Mark = req.Mark
	}
	schedule.UpdateTime = time.Now().Unix()
	err = scheduleService.UpdateSchedule(db.GetDB(), &schedule)
	resp.Schedule = schedule
	if err != nil {
		err = errors.New("修改日程失败")
	}
	return
}

func (scheduleService ScheduleService) DeleteASchedule(ctx context.Context, req metadata.ScheduleRequest) (resp metadata.ScheduleResponse, err error) {
	if req.UserId == 0 {
		err = errors.New("用户不存在")
	}
	var schedule entity.Schedule
	ownerId := ctx.Value(constant.USERID).(int)
	owner, err := scheduleService.GetUserByUserId(db.GetDB(), ownerId)
	if err != nil {
		err = errors.New("用户不存在")
		return
	}

	if owner.ID != req.UserId {
		if owner.Role != constant.UserRoleAdmin {
			err = errors.New("不能修改别人的日程")
			return
		}
	}

	if req.ID == 0 {
		err = errors.New("日程不存在")
		return
	}

	schedule, err = scheduleService.GetScheduleById(db.GetDB(), req.ID)
	if err != nil || schedule.Status == constant.ScheduleStatusDelete || schedule.Status == constant.ScheduleStatusBlock {
		if owner.Role != constant.UserRoleAdmin {
			err = errors.New("日程不存在")
			return
		}
	}

	if schedule.UserId != owner.ID {
		if owner.Role != constant.UserRoleAdmin {
			err = errors.New("不能修改别人的日程")
			return
		}
	}

	schedule.Status = constant.ScheduleStatusDelete
	schedule.UpdateTime = time.Now().Unix()
	err = scheduleService.UpdateSchedule(db.GetDB(), &schedule)

	return

}

func (scheduleService ScheduleService) GetScheduleList(ctx context.Context, req metadata.ScheduleListRequest) (resp metadata.ScheduleListResponse, err error) {

	return
}
