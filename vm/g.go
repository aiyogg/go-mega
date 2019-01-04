package vm

// BaseViewModel view model
type BaseViewModel struct {
	Title       string
	CurrentUser string
}

// SetTitle func - 设置页面 title
func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}

func (v *BaseViewModel) SetCurrentUser(username string) {
	v.CurrentUser = username
}

// BasePageViewModel 分页
type BasePageViewModel struct {
	PrevPage    int
	NextPage    int
	Total       int
	CurrentPage int
	Limit       int // 每页条目数
}

// SetPrevAndNextPage func
func (v *BasePageViewModel) SetPrevAndNextPage() {
	if v.CurrentPage > 1 {
		v.PrevPage = v.CurrentPage - 1
	}

	if (v.Total-1)/v.Limit >= v.CurrentPage {
		v.NextPage = v.CurrentPage + 1
	}
}

// SetBasePageViewModel func
func (v *BasePageViewModel) SetBasePageViewModel(total, page, limit int) {
	v.Total = total
	v.CurrentPage = page
	v.Limit = limit
	v.SetPrevAndNextPage()
}
