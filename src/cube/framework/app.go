package cube

import (
	"errors"
	"net/http"
)

type App struct {
	Router
	Name    string

	server http.Server
	routers []Router
}

//create App
func NewApp(appName string, port int, options map[string]string) (*App, error) {
	app := &App{appName}
	return app
}


//dispatch the request
func (app *App)dispatch(dispatcher Dispatcher) {

}

//global App stack
var apps map[string]App = make(map[string]App)


//get app by the factory mode
func GetApp(appName string, port int, options map[string]string) *App {
	if (appName == nil) {
		errors.New("appName is nil")
	}

	if _, ok := apps[appName]; !ok {
		app, err := NewApp(appName, port, options)
		if err != nil {
			errors.New("GetApp error")
		}
		apps[appName] = app
	}

	return apps[appName]

}