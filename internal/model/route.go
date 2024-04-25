package model

type Route struct {
	Id           int    `json:"id"  gorm:"primarykey" `                    // 路由ID
	Title        string `json:"title"  gorm:"column:title" `               // 路由名称
	Menu_name    string `json:"menu_name"  gorm:"column:menu_name" `       // 菜单名称
	Active_route string `json:"active_route"  gorm:"column:active_route" ` // 动态路由
	Url          string `json:"url"  gorm:"column:url" `                   // 路由地址
	Component    string `json:"component"  gorm:"column:component" `       // 组件地址
	Parent_id    int    `json:"parent_id"  gorm:"column:parent_id" `       // 父级ID
	Description  string `json:"description"  gorm:"column:description" `   // 描述
}

func (m *Route) TableName() string {
	return "pd_route"
}
