//go:build develop

package vite

import (
	"app/pkg/adapter/web"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type ViteProxy struct {
	origin string
}

func NewViteProxy(host string) *ViteProxy {
	return &ViteProxy{
		origin: host,
	}
}

func (p *ViteProxy) ServeProxyData(c *gin.Context) {
	url := url.URL{
		Scheme:   "http",
		Host:     p.origin,
		Path:     c.Request.URL.Path,
		RawQuery: c.Request.URL.RawQuery,
	}
	resp, err := http.Get(url.String())
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Abort()
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		c.Next()
		return
	}

	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	c.Abort()
}

type ViteHTMLRender struct {
	origin   string
	rootPath string
}

func NewViteHTMLRender(origin, rootPath string) *ViteHTMLRender {
	return &ViteHTMLRender{
		origin:   origin,
		rootPath: rootPath,
	}
}

func (r *ViteHTMLRender) Instance(name string, data any) render.Render {
	tmpl, err := r.fetchTemplate(path.Join(r.rootPath, name))
	if err != nil {
		return render.HTML{
			Name: "error",
			Data: data,
		}
	}
	return render.HTML{
		Template: tmpl,
		Name:     path.Join(r.rootPath, name),
		Data:     data,
	}
}

func (r *ViteHTMLRender) fetchTemplate(path string) (*template.Template, error) {
	url := url.URL{
		Scheme: "http",
		Host:   r.origin,
		Path:   path + ".html",
	}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New(path).Parse(string(body))
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func SetViteRouting(server *web.Server, viteProxy *ViteProxy) {
	proxyFunc := func(c *web.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/@") {
			c.Next()
			return
		}
		c.Abort()
		viteProxy.ServeProxyData(c)
	}
	server.Any("/node_modules/", viteProxy.ServeProxyData)
	server.NoRoute(proxyFunc, viteProxy.ServeProxyData)
}
