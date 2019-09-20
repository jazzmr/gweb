package gweb

import (
	"fmt"
	"gweb/conf"
	"gweb/context"
	"log"
	"net/http"
	"sync"
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
			Handler: http.HandlerFunc(nil),
			Pattern: "localhost",
			CxtPool: sync.Pool{
				New: func() interface{} {
					return context.NewContext()
				},
			},
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
