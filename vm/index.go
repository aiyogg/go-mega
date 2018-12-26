package vm

import "github.com/dota2mm/go-mega/model"

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM() IndexViewModel {
	u1, _ := model.GetUserByUsername("Chuck")
	posts, _ := model.GetPostByUserID(u1.ID)

	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *u1, *posts}
	return v
}
