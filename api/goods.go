package api

type CreateGoodsDto struct {
	Name         string `json:"name" `
	Price        int32  `json:"price" 100`
	Source_price int32  `json:"source_price" "10000`
	Stock        int32  `json:"stock" 100`
	Pic          string `json:"pic"`
}
