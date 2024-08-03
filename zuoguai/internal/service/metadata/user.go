package metadata

import "zuoguai/internal/entity"

type UserServiceInterface interface {
	LoginUser(req LoginRequest) (resp LoginResponse, err error)
	RegisterUser(req RegisterRequest) (resp RegisterResponse, err error)
	GetUserList(req GetUserInfoListRequest) (resp UserListResp, err error)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	entity.User
}

type RegisterRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"repeat_password"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
}

type RegisterResponse struct {
	Token string `json:"token"`
	entity.User
}

type UpdateUserInfoRequest struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
	Role           string `json:"role"`
	Phone          string `json:"phone"`
	Avatar         string `json:"avatar"`
	Nickname       string `json:"nickname"`
	Email          string `json:"email"`
	Sex            int    `json:"sex"`
	Age            int    `json:"age"`
	Birthday       int    `json:"birthday"`
	Address        string `json:"address"`
	Remark         string `json:"remark"`
	Status         int    `json:"status"`
}

type GetUserInfoListRequest struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
	Role           string `json:"role"`
	Phone          string `json:"phone"`
	Avatar         string `json:"avatar"`
	Nickname       string `json:"nickname"`
	Email          string `json:"email"`
	Sex            int    `json:"sex"`
	Age            int    `json:"age"`
	Birthday       int    `json:"birthday"`
	Address        string `json:"address"`
	Remark         string `json:"remark"`
	Status         int    `json:"status"`
	PageStart      int    `json:"page_start"`
	PageSize       int    `json:"page_size"`
}

type UserListResp struct {
	Total int64         `json:"tatal"`
	List  []entity.User `json:"list"`
}
