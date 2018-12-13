package vm

// LoginViewModel vm struct
type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

// AddError func for LoginViewModel
func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

// LoginViewModelOp struct - 用于创建 vm 实例
type LoginViewModelOp struct{}

// GetVM func
func (LoginViewModelOp) GetVM() LoginViewModel {
	v := LoginViewModel{} // 利用了 匿名组合 的特性，继承了 BaseViewModel 的 SetTitle 方法
	v.SetTitle("Login")
	return v
}
