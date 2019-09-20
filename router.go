package gweb

import (
	"gweb/conf"
	"gweb/context"
	"log"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

type ControllerInfo struct {
	pattern        string
	controllerType reflect.Type
	methods        map[string]string
	initialize     func() ControllerInterface
}

type ControllerRegister struct {
	Handler http.Handler
	Pattern string
	CxtPool sync.Pool
}

func (c *ControllerRegister) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//c.Handler.ServeHTTP(rw, r)
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
	// method -> controllerInfo
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

	_context := c.CxtPool.Get().(*context.Context)
	_context.Reset(rw, r, reqUri)

	// 重新放回缓存池
	defer c.CxtPool.Put(_context)

	controllerInterface.initCtx(_context)

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

/**
add mappings
path -> ControllerInfo
*/
func Router(path string, c ControllerInterface, mappingMethods ...string) {

	reflectVal := reflect.ValueOf(c)
	t := reflect.Indirect(reflectVal).Type()

	methods := make(map[string]string)
	for _, v := range mappingMethods {
		m := strings.Split(v, ":")
		methods[m[0]] = m[1]
	}

	route := &ControllerInfo{}
	route.controllerType = t
	route.pattern = path
	route.methods = methods
	route.initialize = func() ControllerInterface {
		vc := reflect.New(route.controllerType)
		controllerInterface, ok := vc.Interface().(ControllerInterface)
		if !ok {
			panic("controller is not ControllerInterface")
		}
		return controllerInterface
	}

	gApp.mappings[path] = route
}

/**
find ControllerInfo
*/
func findRouter(path string) (controllerInfo *ControllerInfo, isFind bool) {
	if t, ok := gApp.mappings[path]; ok {
		return t, ok
	}
	return
}
