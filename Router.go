package gweb

import (
	"reflect"
	"strings"
)

type ControllerInfo struct {
	pattern        string
	controllerType reflect.Type
	methods        map[string]string
	initialize     func() ControllerInterface
}

/**
add mappings
path -> Controller
path and http.Method -> Controller.func
*/
func Router(path string, c ControllerInterface, mappingMethods ...string) {
	for _, v := range mappingMethods {
		mappings := make(map[string]reflect.Value)

		if strings.Contains(v, ",") {
			m := strings.Split(v, ",")
			for _, e := range m {
				n := strings.Split(e, ":")
				mappingMethod(c, n, mappings)
			}
		} else {
			n := strings.Split(v, ":")
			mappingMethod(c, n, mappings)
		}
		gApp.methodMappings[path] = mappings
	}
	gApp.mappings[path] = c
}
