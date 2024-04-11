package model

import (
	"time"
)

type Goods struct {
	Id            uint      `json:"id" gorm:"primarykey"`
	Appid         uint      `json:"appid" gorm:"not null"`
	Name          string    `json:"name" gorm:"not null"`
	Mini_name     string    `json:"mini_name"`
	Price         float32   `json:"price" gorm:"not null"`
	Source_price  float32   `json:"source_price"`
	Stock         string    `json:"stock"`
	Pic           string    `json:"pic"`
	Pic_list      string    `json:"pic_list"`
	Quota_num     string    `json:"quota_num"`
	Sales         string    `json:"sales"`
	Virtual_sales string    `json:"virtual_sales"`
	Sort          string    `json:"sort"`
	Views         string    `json:"views"`
	Description   string    `json:"description"`
	Content       string    `json:"content"`
	State         string    `json:"state"`
	Create_time   time.Time `json:"create_time"`
	Update_time   time.Time `json:"update_time"`
}

func (u *Goods) TableName() string {
	return "pd_goods"
}
