package gweb

import (
	"fmt"
	"gweb/conf"
	"gweb/context"
	"gweb/utils"
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
			rw.Write([]byte("找不到页面"))
			return
		}
	}

	mapping := reqUri.Mapping
	//mappingMethod := reqUri.Method
	controllerInfo, ok := findRouter(mapping)
	if !ok {
		// TODO 找不到对应的处理类信息
		rw.Write([]byte("找不到对应的处理类信息!"))
		return
	}

	method, ok := controllerInfo.methods[r.Method]
	if !ok {
		// TODO 找不到对应的处理方法
		rw.Write([]byte("找不到对应的处理方法!"))
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
	//in := param.ConvertParams(methodParams, method.Type(), context)
	out := runMethod.Call(nil)

	log.Println(out)
}

func (c *Controller) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	c.Handler.ServeHTTP(rw, r)
}

func (c *Controller) Input() url.Values {
	if c.Ctx.Request.Form == nil {
		c.Ctx.Request.ParseForm()
	}
	return c.Ctx.Request.Form
}

/**
(path and http.Method) -> func mappings
*/
func mappingMethod(c ControllerInterface, n []string, mappings map[string]reflect.Value) {
	handle := reflect.ValueOf(c)
	handleMethod := handle.MethodByName(utils.UpFirstLetter(n[1]))
	if handleMethod.Kind() == reflect.Func {

		mappings[n[0]] = handleMethod
	} else {
		// TODO error mapping method(path, n[0], n[1])
	}
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
		r.Mapping = "/" + uris[1]
	}

	if len(uris) > 2 {
		r.Method = uris[2]
		if strings.Contains(r.Method, "?") {
			s := strings.Split(r.Method, "?")
			r.Method = s[0]
			rpm := parseParam(s[1])
			r.RequestParams = rpm
		} else {
			if len(uris) > 3 {
				tail := uris[len(uris)-1]
				r.PathParams = uris[3 : len(uris)-1]

				if strings.Contains(tail, "?") {
					s := strings.Split(tail, "?")

					r.PathParams = append(r.PathParams, s[0])

					r.RequestParams = parseParam(s[1])
				} else {
					r.PathParams = uris[3:]
				}
			}
		}
	}

	return r
}

/**
将键值对转化为map对象
*/
func parseParam(p string) map[string]string {
	rp := strings.Split(p, "&")
	rpm := make(map[string]string, 3)
	for _, v := range rp {
		pv := strings.Split(v, "=")
		rpm[pv[0]] = pv[1]
	}
	return rpm
}
