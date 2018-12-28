package vm

import (
	"github.com/dota2mm/go-mega/model"
	"log"
)

type RegisterViewModel struct {
	LoginViewModel
}

type RegisterViewModelOp struct {}
// GetVM method
func (RegisterViewModelOp) GetVM() RegisterViewModel {
	v := RegisterViewModel{}
	v.SetTitle("Register")
	return v
}

// CheckUserExist func - 确认用户是否已经存在
func CheckUserExist(username string) bool {
	_, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Can not find username: ", username)
		return true
	}
	return false
}

// AddUser func
func AddUser(username, password, email string) error {
	return model.AddUser(username, password, email)
}