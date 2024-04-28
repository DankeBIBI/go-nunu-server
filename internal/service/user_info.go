package service

import (
	"context"
	"encoding/json"
	"go-nunu-server/api"
	"go-nunu-server/config"
	"go-nunu-server/internal/model"
	"go-nunu-server/internal/repository"
	"io"
	"net/http"
)

type UserInfoService interface {
	GetUserInfo(ctx context.Context, id int64) (*model.UserInfo, error)

	CheckWechatLogin(info *api.CheckWechatLogin) (interface{}, error)
}

func NewUserInfoService(service *Service, userInfoRepository repository.UserInfoRepository) UserInfoService {
	return &userInfoService{
		Service:            service,
		userInfoRepository: userInfoRepository,
	}
}

type userInfoService struct {
	*Service
	userInfoRepository repository.UserInfoRepository
}

func (s *userInfoService) GetUserInfo(ctx context.Context, id int64) (*model.UserInfo, error) {
	return s.userInfoRepository.GetUserInfo(ctx, id)
}

// 检测是否用微信登录过
func (s *userInfoService) CheckWechatLogin(info *api.CheckWechatLogin) (interface{}, error) {
	type respone struct {
		Openid      string `json:"openid"`
		Session_key string `json:"session_key"`
		Unionid     string `json:"unionid"`
		Errcode     int    `json:"errcode"`
		Errmsg      string `json:"errmsg"`
	}
	params := &api.CheckWechatLogin{
		Code:  info.Code,
		Appid: info.Appid,
	}
	value := config.AppConfig[params.Appid]
	res, err1 := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + value.Appid + "&secret=" + value.Secret + "&js_code=" + params.Code + "&grant_type=authorization_code")
	if err1 != nil {
		return "请求失败", err1
	}
	// 读取响应内容
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "读取响应内容失败:", err
	}
	var data respone
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "解析JSON响应失败", err
	}
	if data.Openid == "" {
		return data.Errmsg, nil
	}
	var userInfo model.UserInfo
	if err := s.userInfoRepository.GetDB().Where("openid = ?", data.Openid).First(&userInfo).Error; err != nil {
		return "未注册", nil
	}
	var user model.User
	if err := s.userInfoRepository.GetDB().Where("u_id = ?", userInfo.U_id).First(&user).Error; err != nil {
		return "用户未注册", nil
	}
	return user, nil
}
