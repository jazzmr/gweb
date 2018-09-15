package gweb

type App struct {
	mappings map[string]ControllerInterface
}

var (
	gApp *App
)

func init() {
	gApp = &App{
		mappings: make(map[string]ControllerInterface),
	}
}

func Add(path string, c ControllerInterface) {
	gApp.mappings[path] = c
}

func getHandler(mapping string) ControllerInterface {
	if mapping == "" {
		return nil
	}
	return gApp.mappings[mapping]
}
