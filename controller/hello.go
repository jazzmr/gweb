package controller

import (
	"fmt"
	"gweb"
)

type HelloController struct {
	gweb.Controller
}

func (h *HelloController) hello() {
	fmt.Println("HelloController")
}
