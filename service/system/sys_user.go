package system

import (
	"errors"
	"fmt"

	"github.com/kinggq/global"
	"github.com/kinggq/model/system"
	"github.com/kinggq/utils"
	"gorm.io/gorm"
)

type UserService struct{}

func (s *UserService) Login(u system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err = global.DB.Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

func (s *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}

	u.Password = utils.BcryptHash(u.Password)
	err = global.DB.Create(&u).Error
	return u, err
}
