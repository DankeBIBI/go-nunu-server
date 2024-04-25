package model

import (
	"time"
)

type Goods struct {
	Id            int       `json:"id" form:"id" gorm:"primarykey"`
	Appid         string    `json:"appid" form:"appid" gorm:"not null"`
	Name          string    `json:"name" form:"name" gorm:"not null"`
	Mini_name     string    `json:"mini_name" form:"mini_name"`
	Price         float32   `json:"price" form:"price" gorm:"not null"`
	Source_price  float32   `json:"source_price" form:"source_price"`
	Stock         int       `json:"stock" form:"stock"`
	Pic           string    `json:"pic" form:"pic"`
	Pic_list      string    `json:"pic_list" form:"pic_list"`
	Quota_num     string    `json:"quota_num" form:"quota_num"`
	Sales         string    `json:"sales" form:"sales"`
	Virtual_sales string    `json:"virtual_sales" form:"virtual_sales"`
	Sort          string    `json:"sort" form:"sort"`
	Views         string    `json:"views" form:"views"`
	Description   string    `json:"description" form:"description"`
	Content       string    `json:"content" form:"content"`
	State         string    `json:"state" form:"state"`
	Create_time   time.Time `json:"create_time" form:"create_time"`
	Update_time   time.Time `json:"update_time" form:"update_time"`
}

func (u *Goods) TableName() string {
	return "pd_goods"
}
