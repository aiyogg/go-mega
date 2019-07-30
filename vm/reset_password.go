package vm

import "github.com/dota2mm/go-mega/model"

type ResetPasswordViewModel struct {
	LoginViewModel
	Token string
}

type ResetPasswordViewModelOp struct{}

// GetVM
func (ResetPasswordViewModelOp) GetVM(token string) ResetPasswordViewModel {
	v := ResetPasswordViewModel{}
	v.SetTitle("Reset Password")
	v.Token = token
	return v
}

// CheckToken 校验token
func CheckToken(tokenString string) (string, error) {
	return model.CheckToken(tokenString)
}

func ResetUserPassword(username, password string) error {
	return model.UpdatePassword(username, password)
}
