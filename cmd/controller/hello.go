package controller

import (
	"fmt"
	"gweb"
)

type HelloController struct {
	gweb.Controller
}

func (c *HelloController) Hello() string {

	//fmt.Println(this.Input().Get("haha"))

	fmt.Println("haha: " + c.Ctx.RequestUri.RequestParams["haha"])

	//c.Ctx.WriteString("hello world")

	return "ret value : hello world"
}

func (c *HelloController) Get() {
	fmt.Println("hello controller Get")
}
