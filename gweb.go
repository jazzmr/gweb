package gweb

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
)

var GRouter *mux.Router

func init() {

	if GRouter == nil {
		GRouter = mux.NewRouter().StrictSlash(true)
	}

	log.Println("init mux.Router success.")
}

func Router(uri string, handler http.Handler, methodName string) {

	t := reflect.TypeOf(handler)

	GRouter.Methods(methodName).Path(uri).Name(t.Name()).Handler(handler)
}
