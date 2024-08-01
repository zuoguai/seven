package metadata

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegistRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
}

type RegistResponse struct {
	Token string `json:"token"`
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
