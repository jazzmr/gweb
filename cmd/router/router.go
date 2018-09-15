package router

import (
	"fmt"
	"gweb"
	"gweb/cmd/controller"
)

func init() {
	gweb.Router("/hello", &controller.HelloController{}, ":Hello", "GET:Hello,DELETE:Delete", "POST:PutHello")
	fmt.Println("register helloController... ...")
}
