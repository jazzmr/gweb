package router

import (
	"fmt"
	"gweb"
	"gweb/cmd/controller"
)

func init() {
	gweb.Router("/hello", &controller.HelloController{}, "GET:Hello")
	fmt.Println("register helloController... ...")
}
