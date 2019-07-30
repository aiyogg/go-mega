package vm

import (
	"github.com/dota2mm/go-mega/config"
	"github.com/dota2mm/go-mega/model"
)

type EmailViewModel struct {
	Username string
	Token    string
	Server   string
}

type EmailViewModelOp struct{}

func (EmailViewModelOp) GetVM(email string) EmailViewModel {
	v := EmailViewModel{}
	u, _ := model.GetUserByEmail(email)
	v.Username = u.Username
	v.Token, _ = u.GenerateToken()
	v.Server = config.GetServerURL()
	return v
}
