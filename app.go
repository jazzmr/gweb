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

func getHandler(mapping string) ControllerInterface {
	if mapping == "" {
		return nil
	}
	return gApp.mappings[mapping]
}
