package vm

import "github.com/dota2mm/go-mega/model"

// ProfileViewModel struct
type ProfileViewModel struct {
	BaseViewModel
	Posts       []model.Post
	Editable    bool
	ProfileUser model.User
}

// ProfileViewModelOp struct
type ProfileViewModelOp struct{}

func (ProfileViewModelOp) GetVM(sUser, pUser string) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u1, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts, _ := model.GetPostByUserID(u1.ID)
	v.ProfileUser = *u1
	v.Editable = (sUser == pUser)
	v.Posts = *posts
	v.SetCurrentUser(sUser)
	return v, nil
}
