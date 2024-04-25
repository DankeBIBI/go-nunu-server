package model

type UserRoute struct {
	Id       int   `json:"id" gorm:"primaryKey"` // 主键id
	Appid    int   `json:"appid" `               // 应用id
	U_id     int   `json:"u_id" `                // 用户id
	Route_id int   `json:"route_id" `            // 路由id
	Route    Route `gorm:"foreignKey:Route_id"`
}

func (m *UserRoute) TableName() string {
	return "pd_user_route"
}
