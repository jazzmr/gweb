package controller

import (
	"fmt"
	"gweb"
)

type HelloController struct {
	gweb.Controller
}

func (h *HelloController) Hello() {
	h.Ctx.WriteString("hello controller Hello")
	fmt.Println("hello controller Hello")
}

func (h *HelloController) Get() {
	fmt.Println("hello controller Get")
}
