package web_test

import (
	"app/pkg/adapter/web"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestExample(t *testing.T) {
	r := web.NewRouter()

	server := web.NewServer()
	server.Use(gin.Logger())

	r.Route("/health", func(g web.IRouter) {
		g.GET("", func(ctx *web.Context) {
			ctx.Status(http.StatusOK)
		})
	})

	r.Route("/api/v1", func(g web.IRouter) {
		g.GET("/profiles", func(ctx *web.Context) {
			ctx.Status(http.StatusOK)
		})
	})

	r.Register(server)
	server.Run()
}
