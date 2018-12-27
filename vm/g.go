package vm

// BaseViewModel view model
type BaseViewModel struct {
	Title string
	CurrentUser string
}

// SetTitle func - 设置页面 title
func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}

func (v *BaseViewModel) SetCurrentUser (username string) {
	v.CurrentUser = username
}