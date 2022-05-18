package structData

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginDataOut struct {
	Token string `json:"token"`
}
