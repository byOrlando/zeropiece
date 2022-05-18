package structData

type RegisterUserData struct {
	Account         string `json:"account" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
	Mobile          string `json:"mobile" binding:"required"`
	Sms             string `json:"sms" binding:"required"`
	Policy          bool   `json:"policy" binding:"required"`
}
