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
path -> ControllerInfo
*/
func Router(path string, c ControllerInterface, mappingMethods ...string) {

	reflectVal := reflect.ValueOf(c)
	t := reflect.Indirect(reflectVal).Type()

	methods := make(map[string]string)
	for _, v := range mappingMethods {
		m := strings.Split(v, ":")
		methods[m[0]] = m[1]
	}

	route := &ControllerInfo{}
	route.controllerType = t
	route.pattern = path
	route.methods = methods
	route.initialize = func() ControllerInterface {
		vc := reflect.New(route.controllerType)
		controllerInterface, ok := vc.Interface().(ControllerInterface)
		if !ok {
			panic("controller is not ControllerInterface")
		}
		return controllerInterface
	}

	gApp.mappings[path] = route
}

func findRouter(path string) (controllerInfo *ControllerInfo, isFind bool) {
	if t, ok := gApp.mappings[path]; ok {
		return t, ok
	}
	return
}
