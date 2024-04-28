package model

import "time"

type ProjectSetting struct {
	Id           int64     `json:"id" gorm:"primaryKey"` // 项目ID
	Appid        int64     `json:"appid"`                // 应用ID
	Project_name string    `json:"project_name"`         // 项目名称
	Index_ad     string    `json:"index_ad"`             // 首页广告
	Index_tip    string    `json:"index_tip"`            // 首页提示
	Create_time  time.Time `json:"create_time"`          // 创建时间
	Update_time  time.Time `json:"update_time"`          // 更新时间
}

func (m *ProjectSetting) TableName() string {
	return "pd_project_setting"
}
