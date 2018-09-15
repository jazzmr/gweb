package controller

import (
	"fmt"
	"gweb"
)

type HelloController struct {
	gweb.Controller
}

func (this *HelloController) Hello() {

	fmt.Println(this.Input().Get("haha"))

	fmt.Println("haha: " + this.Ctx.RequestUri.RequestParams["haha"])

	this.Ctx.WriteString("hello world")
}

func (this *HelloController) Get() {
	fmt.Println("hello controller Get")
}
