package router

import (
	"gweb"
	"gweb/controller"
	"log"
)

func init() {
	gweb.Router("/", &controller.LoginController{}, "Get")
	log.Println("gweb Router init success...")
}
