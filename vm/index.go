package vm

import "github.com/dota2mm/go-mega/model"

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
	Flash string

	BasePageViewModel
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM(username, flash string, page, limit int) IndexViewModel {
	u, _ := model.GetUserByUsername(username)
	posts, total, _ := u.FollowingPostsByPageLimit(page, limit)
	v := IndexViewModel{}
	v.SetTitle("Homepage")
	v.Posts = *posts
	v.Flash = flash
	v.SetBasePageViewModel(total, page, limit)
	v.SetCurrentUser(username)
	return v
}

// CreatePost 创建文章
func CreatePost(username, post string) error {
	u, _ := model.GetUserByUsername(username)
	return u.CreatePost(post)
}
