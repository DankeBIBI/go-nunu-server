package api

type AppConfig struct {
	Appid int `json:"appid"`
}
type CreateGoodsDto struct {
	AppConfig
	Name         string `json:"name" `
	Price        int    `json:"price"`
	Source_price int    `json:"source_price" `
	Stock        int    `json:"stock" `
	Pic          string `json:"pic"`
}
