package gweb

import (
	"github.com/gorilla/mux"
	"log"
)

type ControllerInterface interface {
	Get()
	Post()
	Put()
	Delete()
}

type Controller struct {
	Data          map[interface{}]interface{}
	Name          string
	MethodMapping map[string]string
}

var router *mux.Router

func init() {
	if router == nil {
		router = mux.NewRouter().StrictSlash(true)
	}

	log.Println("init mux.Router success.")
}

func (c *Controller) Get() {
	log.Println("get... ...")
}

func (c *Controller) Post() {

}

func (c *Controller) Put() {

}

func (c *Controller) Delete() {

}

func Router(uri string, c ControllerInterface, methodName string) {
	router.Methods(methodName).Path(uri).Name("").Handler(nil)
}
