package router

import (
	"gweb"
	"gweb/controller"
)

func init() {
	gweb.Router("/", &controller.LoginController{}, "")
}
