package service

import (
	"errors"
	"fmt"
	"go-nunu-server/api"
	"go-nunu-server/internal/model"
	"go-nunu-server/internal/repository"
)

type UserService interface {
	GetUserList(phone string) (interface{}, error)
	Login(info *api.LoginDto) (interface{}, error)
}

func NewUserService(service *Service, userRepository repository.UserRepository) UserService {
	return &userService{
		Service:        service,
		userRepository: userRepository,
	}
}

type userService struct {
	*Service
	userRepository     repository.UserRepository
	userInfoRepository repository.UserInfoRepository
}

// 获取用户列表
func (s *userService) GetUserList(phone string) (interface{}, error) {
	if phone != "" {
		var user model.User
		if err := s.userRepository.GetDB().Find(&user, "phone = ?", phone).Error; err != nil {
			return nil, err
		}
		return user, nil
	}
	var user []model.User
	if err := s.userRepository.GetDB().Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) checkUser(userId interface{}) (model.User, error) {
	var user model.User
	noRegist := false
	//没有注册
	if err := s.userRepository.GetDB().Where("u_id = ?", userId).First(&user).Error; err != nil {
		noRegist = true
	}
	if err := s.userRepository.GetDB().Where("phone = ?", userId).First(&user).Error; err != nil {
		noRegist = true
	}
	fmt.Println(user)
	if noRegist {
		return user, errors.New("用户不存在")
	}
	return user, nil
}

// 用户登录
func (s *userService) Login(info *api.LoginDto) (interface{}, error) {
	value, err := s.checkUser(info.User)
	fmt.Println(value)
	if err != nil {
		return "用户不存在", err
	}

	var route []model.UserRoute
	if err := s.userRepository.GetDB().Where("u_id =?", info.User).Preload("Route").Find(&route).Error; err != nil {
		value.UserRoute = nil
	}
	value.UserRoute = route
	fmt.Println(value)
	return value, nil
}
