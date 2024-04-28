package model

type App struct {
	Id           int            `json:"id" form:"id" gorm:"primaryKey"`                // 项目ID
	AppName      string         `json:"appName" form:"appName" gorm:"column:app_name"` // 项目名称
	Logo         string         `json:"logo" form:"logo" gorm:"column:logo"`           // 项目logo
	Project_name string         `json:"project_name"`                                  // 项目名称
	Index_ad     string         `json:"index_ad"`                                      // 首页广告
	Index_tip    string         `json:"index_tip"`                                     // 首页提示
	Theme        []ProjectTheme `json:"theme" gorm:"foreignKey:Appid;references:Id;"`
}

func (m *App) TableName() string {
	return "pd_app"
}
