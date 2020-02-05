package auth_service

import "errors"

type Auth struct {
	Username string
	Password string
}

//func (a *Auth) Check() (bool, error) {
//	return models.CheckAuth(a.Username, a.Password)
//}

func (a *Auth) Check() (bool, error) {
	if a.Username == "root" && a.Password == "123456" {
		return true, nil
	}
	error := errors.New("认证信息错误")
	return false, error
}
