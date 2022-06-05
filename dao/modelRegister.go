package dao

func RegistrationModel() []interface{} {
	return []interface{}{
		&CityCode{},
		&PeekabooUser{},
		&Role{},
		&DouyinUser{},
		&DouyinWorks{},
		&CityCode{},
		&InsideWeChatOpenID{},
	}

}
