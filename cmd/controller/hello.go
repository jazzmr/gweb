package controller

import (
	"fmt"
	"gweb"
)

type HelloController struct {
	gweb.Controller
}

func (c *HelloController) Hello() string {

	fmt.Println(c.Input().Get("a"))
	fmt.Println(c.Input())

	//fmt.Println("haha: " + c.Ctx.RequestUri.RequestParams["haha"])

	//c.Ctx.WriteString("hello world")

	return "<html> <head> </head> <body> <font color='red'> ret value : hello world </font> </body></html>"
}

func (c *HelloController) Get() {
	fmt.Println("hello controller Get")
}
