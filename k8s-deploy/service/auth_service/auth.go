package auth_service

type Auth struct {
	Username string
	Password string
}

//func (a *Auth) Check() (bool, error) {
//	return models.CheckAuth(a.Username, a.Password)
//}

func (a *Auth) Check() (bool, error) {
	if a.Username != "root" && a.Password != nil() {
		return false, nil
	}

}
