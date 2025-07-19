package web

import (
	"iter"

	"github.com/gin-gonic/gin"
)

type IRouter = gin.IRouter

type Router struct {
	m     map[string]RegisterFunc
	order []string
}

func NewRouter() *Router {
	return &Router{}
}

type RegisterFunc func(IRouter)

func (r *Router) Route(path string, f RegisterFunc) {
	if r.m == nil {
		r.m = make(map[string]RegisterFunc)
	}
	if _, ok := r.m[path]; ok {
		panic("already exist path")
	}
	r.m[path] = f
	r.order = append(r.order, path)
}

func (r *Router) Register(server *Server) {
	for path, f := range r.routes() {
		g := server.Group(path)
		f(g)
	}
}

func (r *Router) routes() iter.Seq2[string, RegisterFunc] {
	return func(yield func(path string, f RegisterFunc) bool) {
		for _, _path := range r.order {
			if !yield(_path, r.m[_path]) {
				return
			}
		}
	}
}
