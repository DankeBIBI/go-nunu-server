package config

type App struct {
	App_name   string
	Appid      string
	TemplateID []string
	Secret     string
}

var AppConfig = map[string]App{
	"1": {
		App_name: "蛋壳模拟器",
		Appid:    "wxa95a278493f56132",
		TemplateID: []string{
			"RkF_F2vMmSp0TcxzEd19ernPdTpUWKYV8mcEZGRGGi0",
		},
		Secret: "e1a9cd092c7a025099b13e48f5b5d6b7",
	},
}

// 小程序配置
func Apps(appid string) App {
	return AppConfig[appid]
}
