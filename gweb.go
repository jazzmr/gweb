package gweb

import (
	"fmt"
	"gweb/conf"
	"gweb/context"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type ControllerInterface interface {
	Get()
	Post()
	Put()
	Delete()
	Init(ctx *context.Context)
}

type Controller struct {
	Handler http.Handler
	Pattern string
	Ctx     *context.Context
}

func (c *Controller) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	c.Handler.ServeHTTP(rw, r)
}

func Run() {
	config := conf.GetConfig()

	time.Sleep(1 * time.Second)

	fmt.Println(config.Server)

	h := &Controller{
		Handler: http.HandlerFunc(dispatch),
		Pattern: "localhost",
	}

	log.Println("gweb start success ... ...")
	e := http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), h)
	log.Print("e : ", e)
}

/**
  call this method
*/
func dispatch(rw http.ResponseWriter, r *http.Request) {

	uri := r.RequestURI

	reqUri := parseURI(uri)

	p := conf.GetContextPath()

	if p != "" {
		if p != reqUri.ContextPath {
			rw.WriteHeader(404)
			context.WriterString(rw, "找不到页面!")
			return
		}
	}

	mapping := reqUri.Mapping
	//mappingMethod := reqUri.Method
	controllerInfo, ok := findRouter(mapping)
	if !ok {
		context.WriterString(rw, "找不到对应的处理类信息!")
		return
	}

	method, ok := controllerInfo.methods[r.Method]
	if !ok {
		context.WriterString(rw, "找不到对应的处理方法!")
		return
	}

	controllerInterface := controllerInfo.initialize()

	_context := &context.Context{
		Request:        r,
		ResponseWriter: rw,
		RequestUri:     reqUri,
	}
	controllerInterface.Init(_context)

	vc := reflect.ValueOf(controllerInterface)
	runMethod := vc.MethodByName(method)
	ret := runMethod.Call(nil)

	var retString string
	for _, v := range ret {
		retString += v.Interface().(string)
	}

	context.WriterString(rw, retString)

	log.Println("ret : ", ret)
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
