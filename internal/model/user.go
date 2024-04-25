package model

import (
	"time"
)

type User struct {
	Id          int         `json:"id" gorm:"primarykey"`                               // 主键ID
	Appid       int         `json:"appid"`                                              // 应用ID
	U_id        int         `json:"u_id"`                                               // 用户ID
	U_name      string      `json:"u_name"`                                             // 用户名
	Phone       string      `json:"phone"`                                              // 手机号
	Head_url    string      `json:"head_url"`                                           // 头像地址
	Sex         int         `json:"sex"`                                                // 性别
	Age         int         `json:"age"`                                                // 年龄
	Integral    int         `json:"integral"`                                           // 积分
	Level       int         `json:"level"`                                              // 等级
	Create_time time.Time   `json:"create_time" gorm:"column:CreatedAt;type:datetime;"` // 创建时间
	Update_time time.Time   `json:"update_time" gorm:"column:UpdatedAt;type:datetime;"` // 更新时间
	UserRoute   []UserRoute `gorm:"foreignKey:U_id;references:Id;"`                     // 用户路由
}

func (m *User) TableName() string {
	return "pd_user"
}
