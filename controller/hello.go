package controller

import (
	"fmt"
	"gweb"
)

type HelloController struct {
	gweb.Controller
}

func (h *HelloController) Hello() {
	h.Ctx.WriteString("hello world")
}

func (h *HelloController) Get() {
	fmt.Println("hello controller Get")
}
