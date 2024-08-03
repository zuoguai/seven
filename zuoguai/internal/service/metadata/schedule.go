package metadata

import (
	"golang.org/x/net/context"
	"zuoguai/internal/entity"
)

type ScheduleServiceInterface interface {
	GetASchedule(context context.Context, req ScheduleRequest) (resp ScheduleResponse, err error)
	AddASchedule(context context.Context, req ScheduleRequest) (resp ScheduleResponse, err error)
	UpdateASchedule(context context.Context, req ScheduleRequest) (resp ScheduleResponse, err error)
	DeleteASchedule(context context.Context, req ScheduleRequest) (resp ScheduleResponse, err error)
	GetScheduleList(context context.Context, req ScheduleListRequest) (resp ScheduleListResponse, err error)
}

type ScheduleRequest struct {
	entity.Schedule
}

type ScheduleResponse struct {
	entity.Schedule
}

type ScheduleListRequest struct {
	entity.Schedule
	PageStart int `json:"page_start"`
	PageSize  int `json:"page_size"`
}

type ScheduleListResponse struct {
	Count int               `json:"count"`
	List  []entity.Schedule `json:"list"`
}
