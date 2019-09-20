package router

import (
	"fmt"
	"gweb"
	"gweb/cmd/controller"
)

func init() {
	gweb.Router("/hello", &controller.HelloController{}, "GET:Hello")
	gweb.Router("/login", &controller.LoginController{}, "GET:Login")
	fmt.Println("register helloController... ...")
}
