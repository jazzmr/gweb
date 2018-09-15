package gweb

import "reflect"

type App struct {
	mappings       map[string]ControllerInterface
	methodMappings map[string]map[string]reflect.Value
}

var (
	gApp *App
)

func init() {
	gApp = &App{
		mappings:       make(map[string]ControllerInterface),
		methodMappings: make(map[string]map[string]reflect.Value),
	}
}

func getHandleMethod(path, httpMethod string) reflect.Value {
	if path == "" {
		return reflect.Value{}
	}

	if v, ok := gApp.methodMappings[path]; ok {
		return v[httpMethod]
	}
	return reflect.Value{}
}
