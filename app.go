package gweb

import (
	"fmt"
	"gweb/conf"
	"log"
	"net/http"
)

type App struct {
	Handler  *ControllerRegister
	mappings map[string]*ControllerInfo
}

var (
	gApp *App
)

func init() {
	gApp = &App{
		Handler: &ControllerRegister{
			Handler: http.HandlerFunc(dispatch),
			Pattern: "localhost",
		},
		mappings: make(map[string]*ControllerInfo),
	}
}

func Run() {
	config := conf.GetConfig()

	log.Println(config.Server)
	log.Println("gweb start success ... ...")

	e := http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), gApp.Handler)
	log.Print("e : ", e)
}
