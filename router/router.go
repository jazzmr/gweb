package router

import (
	"gweb"
	"gweb/controller"
	"log"
	"reflect"
)

func init() {
	gweb.Router("/test", &controller.LoginController{}, "Get")

	h := &controller.HelloController{}

	t := reflect.TypeOf(h)
	mLen := t.NumMethod()

	for i := 0; i < mLen; i++ {

		m := t.Method(i)

		if m.Name != "ServeHTTP" {
			gweb.Router(m.Name, h, h.Method)
		}
	}

	log.Println("gweb Router init success...")
}
