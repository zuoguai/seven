package user

import (
	"errors"
	"fmt"
	"time"
	"zuoguai/internal/db"
	"zuoguai/internal/entity"
	"zuoguai/internal/service/metadata"
	"zuoguai/pkg/constant"
	"zuoguai/utils"
)

func (userService UserService) LoginUser(req metadata.LoginRequest) (resp metadata.LoginResponse, err error) {
	//
	var user entity.User
	var token string
	if req.Password == "" || req.Username == "" {
		err = errors.New("请填写完整的用户名和密码")
		return
	}
	if user, err = entity.GetEntities().UserEntity.GetUserByUsername(db.GetDB(), req.Username); err != nil {
		err = errors.New("用户名或密码有误")
		return
	}
	if user.Password != req.Password {
		err = errors.New("用户名或密码有误")
		return
	}
	if token, err = utils.GenerateJwtToken("", "作怪", user.ID, -1); err != nil {
		err = errors.New("生成token失败")
		return
	}
	fmt.Println(user, " 登录")

	user.Password = "******"
	user.DeleteAt = 0
	user.Status = 0
	resp = metadata.LoginResponse{
		Token: token,
		User:  user,
	}

	return
}

func (userService UserService) RegisterUser(req metadata.RegisterRequest) (resp metadata.RegisterResponse, err error) {
	//
	var user entity.User
	var token string
	if req.Password == "" || req.Username == "" || req.PasswordRepeat == "" {
		err = errors.New("请填写完整的用户名和密码")
		return
	}
	if len(req.Username) < 2 {
		err = errors.New("用户名过短")
		return
	}
	if len(req.Username) > 20 {
		err = errors.New("用户名过长")
		return
	}
	if len(req.Password) < 6 {
		err = errors.New("密码过短")
		return
	}
	if len(req.Username) > 20 {
		err = errors.New("密码过短")
		return
	}
	if req.Password != req.PasswordRepeat {
		err = errors.New("两次密码输入不一致")
		return
	}
	if user, err = entity.GetEntities().UserEntity.GetUserByUsername(db.GetDB(), req.Username); err == nil {
		fmt.Println(user)
		err = errors.New("用户名已存在")
		return
	}

	user.Username = req.Username
	user.Password = req.Password
	user.Email = req.Email
	user.Phone = req.Phone
	user.Status = constant.UserStatusNormal
	user.CreateAt = time.Now().Unix()
	user.UpdateAt = user.CreateAt
	user.Remark = "天生我材必有用"
	user.Role = constant.UserRoleNormal
	user.Nickname = "天下无双"
	user.Birthday = user.CreateAt

	if err = entity.GetEntities().UserEntity.CreateUser(db.GetDB(), &user); err != nil {
		err = errors.New("服务器异常，注册失败")
		return
	}

	if token, err = utils.GenerateJwtToken("", "作怪", user.ID, -1); err != nil {
		fmt.Println("生成token失败")
		err = errors.New("生成token失败")
		return
	}

	user.Password = "******"
	resp = metadata.RegisterResponse{
		Token: token,
		User:  user,
	}
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
