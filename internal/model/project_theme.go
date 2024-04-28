package model

type ProjectTheme struct {
	Id         int    `json:"id" form:"id" gorm:"primaryKey" `                       // 项目主题ID
	Appid      int    `json:"appid" form:"appid" gorm:"column:appid"`                // 项目ID
	Theme_Name string `json:"theme_name" form:"theme_name" gorm:"column:theme_name"` // 项目主题名称
	Bg_main    string `json:"--bg_main" form:"bg_main" gorm:"column:bg_main"`        // 项目主题背景1
	Bg_sec     string `json:"--bg_sec" form:"bg_sec" gorm:"column:bg_sec"`           // 项目主题背景2
	Bg_last    string `json:"--bg_last" form:"bg_last" gorm:"column:bg_last"`        // 项目主题背景3
	Tc_main    string `json:"--tc_main" form:"tc_main" gorm:"column:tc_main"`        // 项目主题文字1
	Tc_sec     string `json:"--tc_sec" form:"tc_sec" gorm:"column:tc_sec"`           // 项目主题文字2
	Tc_last    string `json:"--tc_last" form:"tc_last" gorm:"column:tc_last"`        // 项目主题文字3
}

func (m *ProjectTheme) TableName() string {
	return "pd_project_theme"
}
