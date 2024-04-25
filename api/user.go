package api

// 登录
type LoginDto struct {
	AppConfig
	User     string `json:"user" form:"user"`
	Password string `json:"password" form:"password"`
}
