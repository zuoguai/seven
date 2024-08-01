package user

import (
	"zuoguai/internal/db"
	"zuoguai/internal/entity"
	"zuoguai/internal/service/metadata"
)

func (userService UserService) LoginUser(req metadata.LoginRequest) (resp metadata.LoginRequest, err error) {
	//

	return
}

func (userService UserService) RegistUser(req metadata.RegistRequest) (resp metadata.RegistResponse, err error) {
	//

	return
}

func (userService UserService) GetUserList(req metadata.GetUserInfoListRequest) (resp metadata.UserListResp, err error) {
	//

	var options []entity.UserFindOptionFn

	if req.ID != 0 {
		options = append(options, entity.UserFindOptionIns.WithID(req.ID))
	}

	if req.Username != "" {
		options = append(options, entity.UserFindOptionIns.WithUsername(req.Username))
	}

	if req.Role != "" {
		options = append(options, entity.UserFindOptionIns.WithRole(req.Role))
	}
	if req.Email != "" {
		options = append(options, entity.UserFindOptionIns.WithEmail(req.Email))
	}
	if req.Phone != "" {
		options = append(options, entity.UserFindOptionIns.WithPhone(req.Phone))
	}
	if req.Address != "" {
		options = append(options, entity.UserFindOptionIns.WithAddress(req.Address))
	}
	if req.Status != 0 {
		options = append(options, entity.UserFindOptionIns.WithStatus(req.Status))
	}
	if req.Sex != 0 {
		options = append(options, entity.UserFindOptionIns.WithSex(req.Sex))
	}

	resp.Total, err = userService.userEntity.FindUserCountByCond(db.GetDB(), options...)
	if err != nil {
		return
	}

	if req.PageStart > 0 && req.PageSize > 0 {
		options = append(options, entity.UserFindOptionIns.WithPage((req.PageStart-1)*req.PageSize, req.PageSize))
	} else if req.PageSize > 0 {
		options = append(options, entity.UserFindOptionIns.WithPage(0, req.PageSize))
	}

	resp.List, err = userService.userEntity.FindUserListByCond(db.GetDB(), options...)

	return resp, err
}
