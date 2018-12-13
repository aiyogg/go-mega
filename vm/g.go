package vm

// BaseViewModel view model
type BaseViewModel struct {
	Title string
}

// SetTitle func - 设置页面 title
func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}
