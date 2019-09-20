package gweb

import (
	"gweb/context"
	"net/url"
	"strings"
)

type ControllerInterface interface {
	Get()
	Post()
	Put()
	Delete()
	Init(ctx *context.Context)
}

type Controller struct {
	Ctx *context.Context
}

func (c *Controller) Input() url.Values {
	if c.Ctx.Request.Form == nil {
		c.Ctx.Request.ParseForm()
	}
	return c.Ctx.Request.Form
}

func (c *Controller) Get() {
}

func (c *Controller) Post() {

}

func (c *Controller) Put() {

}

func (c *Controller) Delete() {

}

func (c *Controller) Init(ctx *context.Context) {
	c.Ctx = ctx
}

func parseURI(uri string) *context.RequestUri {
	uris := strings.Split(uri, "/")

	uris = uris[1:]
	r := &context.RequestUri{}

	r.ContextPath = "/" + uris[0]

	if len(uris) > 1 {
		if strings.Contains(uri, "?") {
			lastElement := uris[len(uris)-1]
			strings.Split(lastElement, "?")
			r.Mapping = "/" + strings.Join(uris[1:len(uris)-1], "/") + strings.Split(lastElement, "?")[0]
		} else {
			r.Mapping = "/" + strings.Join(uris[1:], "/")
		}
	}
	return r
}
