package vm

import (
	"github.com/dota2mm/go-mega/model"
	"log"
)

// ResetPasswordRequestViewModel struct
type ResetPasswordRequestViewModel struct {
	LoginViewModel
}

// ResetPasswordRequestViewModelOp struct
type ResetPasswordRequestViewModelOp struct{}

// GetVM func
func (*ResetPasswordRequestViewModelOp) GetVM() ResetPasswordRequestViewModel {
	v := ResetPasswordRequestViewModel{}
	v.SetTitle("Forgot Password")
	return v
}

// CheckEmailExist 确认邮箱是否存在
func CheckEmailExist(email string) bool {
	_, err := model.GetUserByEmail(email)
	if err != nil {
		log.Panicln("Can not find email", email)
		return false
	}
	return true
}
