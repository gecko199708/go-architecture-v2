package logic

import "sync/atomic"

type dependencies struct {
	userLogics  atomic.Pointer[UserLogics]
	salesLogics atomic.Pointer[SalesLogics]
}

func newDependencies() *dependencies {
	return &dependencies{}
}

func (d *dependencies) SetUserLogics(logics UserLogics) {
	d.userLogics.Store(&logics)
}

func (d *dependencies) SetSalesLogics(logics SalesLogics) {
	d.salesLogics.Store(&logics)
}
