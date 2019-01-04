package vm

import "github.com/dota2mm/go-mega/model"

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
	Flash string
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM(username, flash string) IndexViewModel {
	u, _ := model.GetUserByUsername(username)
	posts, _ := u.FollowingPosts()

	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *posts, flash}
	v.SetCurrentUser(username)
	return v
}

// CreatePost 创建文章
func CreatePost(username, post string) error {
	u, _ := model.GetUserByUsername(username)
	return u.CreatePost(post)
}
