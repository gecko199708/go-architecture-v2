package logic

import "sync/atomic"

var deps atomic.Pointer[dependencies]

func Dependencies() *dependencies {
	if v := deps.Load(); v != nil {
		return v
	}
	deps.Store(newDependencies())
	return deps.Load()
}

func Users() UserLogics {
	if v := Dependencies().userLogics.Load(); v != nil {
		return v
	}
	return nil
}

func Sales() SalesLogics {
	if v := Dependencies().salesLogics.Load(); v != nil {
		return v
	}
	return nil
}
