package web

import (
	"net/http"
	"cube/utils"
	"errors"
)

//middleware is function
type MiddleWare func(req Request, res Response, next func())


//router dispatcher
type Dispatcher struct {
	handlers    utils.List
	middleWares utils.List
	app         *App
}


//base router will dispatcher
func (dispatcher *Dispatcher)ServeHTTP(writer http.ResponseWriter, r *http.Request) {

	router := NewRouter(r)

	dispatcher.Dispatch(router)

}

//add the router middleware
//mode1: Use("/pattern/",MiddleWare)
//mode2: Use(MiddleWare)
func (dispatcher *Dispatcher)Use(args...interface{}) error {
	switch len(args) {
	case 1:
		dispatcher.middleWares.Push(args[0])
	case 2:
		m := map[string]MiddleWare{}
		m[args[0]] = args[1]
		dispatcher.handlers.Push(m)
	default:
		errors.New("Use MiddleWare Params Error")
	}

}

//exec router dispatch
func (dispatcher *Dispatcher)Dispatch(router Router) {

}

func routerMatch() {

}

func routerNext() {

}