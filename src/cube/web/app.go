package web

import (
	"errors"
	"net/http"
	"time"
	"log"
	"strconv"
)


//Web App Unit Instance
//It can support the service with one port
type App struct {
	Name       string       //appName

	dispatcher *Dispatcher
	server     *http.Server //http.Server instance
}


//create App
func NewApp(appName string, port int) *App {
	dispatcher := &Dispatcher{
	}

	s := &http.Server{
		Addr:           ":" + strconv.Itoa(port),
		Handler:        dispatcher,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	app := &App{Name:appName, server:s, dispatcher:dispatcher}

	dispatcher.app = app

	return app
}

//start app ListenAndServe
func (app *App)Start() {
	err := app.server.ListenAndServe()
	if (err != nil) {
		log.Fatalln(err)
	}
}


//global App stack
var apps map[string]*App = make(map[string]*App)
//get app by the factory mode
func GetApp(appName string, port int) *App {
	if (appName == "") {
		errors.New("appName is nil")
	}

	if _, ok := apps[appName]; !ok {
		apps[appName] = NewApp(appName, port)
	}

	return apps[appName]
}