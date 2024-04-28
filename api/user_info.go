package api
type CheckWechatLogin struct {
	Code  string `json:"code" form:"code"`
	Appid string `json:"appid" form:"appid"`
}