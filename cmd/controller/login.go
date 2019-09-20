package controller

import (
	"fmt"
	"gweb"
	"strings"
)

type LoginController struct {
	gweb.Controller
}

func (c *LoginController) Login() {

	a := c.Input().Get("a")

	if strings.Contains(a, "a") {
		goto testGoto
	}
	fmt.Println("test")

testGoto:
	fmt.Println("testGoto")
}
