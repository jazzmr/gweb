package gweb

import (
	"fmt"
	"gweb/conf"
	"gweb/context"
	"gweb/utils"
	"log"
	"net/http"
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

type requestURI struct {
	contextPath   string
	mapping       string
	method        string
	pathParams    []string
	requestParams map[string]string
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
		if p != reqUri.contextPath {
			rw.WriteHeader(404)
			rw.Write([]byte("找不到页面"))
			return
		}
	}

	mapping = reqUri.mapping
	mappingMethod = reqUri.method

	ci := getHandler(mapping)

	// handler is nil, give a default handler
	if ci == nil {
		rw.Write([]byte("there is no mapping handler found."))
		return
	}

	ci.Init(&context.Context{ResponseWriter: rw})

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

func Router(path string, c ControllerInterface, mappingMethods ...string) {
	Add(path, c)
}

func parseURI(uri string) *requestURI {
	uris := strings.Split(uri, "/")

	uris = uris[1:]
	r := &requestURI{}

	r.contextPath = "/" + uris[0]

	if len(uris) > 1 {
		r.mapping = "/" + uris[1]
	}

	if len(uris) > 2 {
		r.method = uris[2]
	}

	if len(uris) > 3 {
		tail := uris[len(uris)-1]
		if strings.Contains(tail, "?") {
			r.pathParams = uris[3 : len(uris)-1]
			rp := strings.Split(tail, "&")
			rpm := make(map[string]string, 3)
			for _, v := range rp {
				pv := strings.Split(v, "=")
				rpm[pv[0]] = pv[1]
			}

			r.requestParams = rpm
		} else {
			r.pathParams = uris[3:]
		}
	}

	return r
}
