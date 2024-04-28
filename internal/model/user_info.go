package model

type UserInfo struct {
	Id         int    `json:"id" gorm:"primaryKey"` // ID
	Appid      int    `json:"appid"`                // 应用ID
	U_id       int    `json:"u_id"`                 // 用户ID
	U_password string `json:"u_password"`           // 用户密码
	Openid     string `json:"openid"`               // 微信openid

}

func (m *UserInfo) TableName() string {
	return "pd_user_info"
}
