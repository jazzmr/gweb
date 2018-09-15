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
	conf := conf.GetConfig()

	time.Sleep(1 * time.Second)

	fmt.Println(conf.Server)

	h := &Controller{
		Handler: http.HandlerFunc(dispatch),
		Pattern: "localhost",
	}

	log.Println("gweb start success ... ...")
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Server.Port), h)
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

func Router(path string, c ControllerInterface, mappingMethods ...string) {
	for _, v := range mappingMethods {
		mappings := make(map[string]reflect.Value)

		if strings.Contains(v, ",") {
			m := strings.Split(v, ",")
			for _, e := range m {
				n := strings.Split(e, ":")

				mappingMethod(c, n, mappings)
			}
		} else {
			n := strings.Split(v, ":")
			mappingMethod(c, n, mappings)
		}

		gApp.methodMappings[path] = mappings
	}
	gApp.mappings[path] = c
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

/**
  call this method
*/
func dispatch(rw http.ResponseWriter, r *http.Request) {

	var mapping, mappingMethod string

	//m := r.Method

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

	mapping = reqUri.Mapping
	mappingMethod = reqUri.Method

	ci := getHandler(mapping)

	// handler is nil, give a default handler
	if ci == nil {
		rw.Write([]byte("there is no mapping handler found."))
		return
	}

	ci.Init(&context.Context{
		ResponseWriter: rw,
		Request:        r,
		RequestUri:     reqUri,
	})

	getValue := reflect.ValueOf(ci)

	// 一定要指定参数为正确的方法名
	// 2. 先看看带有参数的调用方法
	//methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
	//args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
	//methodValue.Call(args)

	// 一定要指定参数为正确的方法名
	// 3. 再看看无参数的调用方法
	methodValue := getValue.MethodByName(utils.UpFirstLetter(mappingMethod))

	if methodValue.Kind() != reflect.Func {
		rw.Write([]byte("there is no mapping method found."))
		return
	}

	args := make([]reflect.Value, 0)
	methodValue.Call(args)

	//fmt.Printf("hello world, %s, %s", m, uri)
	//rw.Write([]byte("hello world!"))
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

func parseParam(p string) map[string]string {
	rp := strings.Split(p, "&")
	rpm := make(map[string]string, 3)
	for _, v := range rp {
		pv := strings.Split(v, "=")
		rpm[pv[0]] = pv[1]
	}
	return rpm
}
