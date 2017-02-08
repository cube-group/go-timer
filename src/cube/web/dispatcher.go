package web

import (
	"net/http"
	"cube/utils"
	"errors"
)

//middleware is function
type MiddleWare struct {
	Pattern  string
	Instance func(req Request, res Response, next func())
}


//router dispatcher
type Dispatcher struct {
	router      Router
	middleWares utils.List

	app         *App
}


//base router will dispatcher
func (dispatcher *Dispatcher)ServeHTTP(writer http.ResponseWriter, r *http.Request) {

	router := NewRouter(writer, r)

	dispatcher.Dispatch(router)

}

//add the router middleware
//mode1: Use("/pattern/",MiddleWare)
//mode2: Use(MiddleWare)
func (dispatcher *Dispatcher)Use(args...interface{}) error {
	var err error

	switch len(args) {
	case 1:
		dispatcher.middleWares.Push(MiddleWare{"MiddleWare", args[0]})
	case 2:
		dispatcher.middleWares.Push(MiddleWare{args[0], args[1]})
	default:
		err = errors.New("Use MiddleWare Params Error")
	}
	return err
}

//exec router dispatch
func (dispatcher *Dispatcher)Dispatch(router Router) {
	dispatcher.router = router

	dispatcher.routerNext()
}

func (dispatcher *Dispatcher)routerNext() {
	if mw := dispatcher.middleWares.Current(); mw {
		middleWare := mw.(MiddleWare)
		if middleWare.Pattern == "MiddleWare" {
			middleWare.Instance(dispatcher.router.req, dispatcher.router.res, dispatcher.routerNext)
		}else {
			if routerMatch(dispatcher.router.pattern, middleWare.Pattern) {
				middleWare.Instance(dispatcher.router.req, dispatcher.router.res, dispatcher.routerNext)
			}else {
				dispatcher.routerNext()
			}
		}
	}else {
		//404
	}
}

func routerMatch(urlString string, routerPattern string) bool {

}

