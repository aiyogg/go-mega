package vm

import "github.com/dota2mm/go-mega/model"

type ExploreViewModel struct {
	BaseViewModel
	Posts []model.Post
	BasePageViewModel
}

type ExploreViewModelOp struct {}

func (ExploreViewModelOp) GetVM(username string, page, limit int) (ExploreViewModel, error) {
	v := ExploreViewModel{}
	v.SetTitle("Explore")

	posts, total, err := model.GetPostsByPageAndLimit(page, limit)
	if err != nil {
		return v, nil
	}
	v.Posts = *posts
	v.SetBasePageViewModel(total, page, limit)
	v.SetCurrentUser(username)
	return v, nil
}
