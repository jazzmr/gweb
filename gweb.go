package gweb

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
)

var MRouter *mux.Router

func init() {

	if MRouter == nil {
		MRouter = mux.NewRouter().StrictSlash(true)
	}

	log.Println("init mux.Router success.")
}

func Router(uri string, handler http.Handler, methodName string) {

	t := reflect.TypeOf(handler)

	MRouter.Methods(methodName).Path(uri).Name(t.Name()).Handler(handler)
}
