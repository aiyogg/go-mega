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
	u1 := model.User{Username: "Chuck"}
	u2 := model.User{Username: "Jay"}

	posts := []model.Post{
		model.Post{User: u1, Body: "Chuck is look like so cool!"},
		model.Post{User: u2, Body: "Jay is best cool boy!"},
	}

	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, u1, posts}
	return v
}
