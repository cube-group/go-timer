package cube

type IRouter interface {
	use(pattern string, MiddleWare)
}

//middleware is function
type MiddleWare func(req Request, res Response, next func())

//core Router struct
type Router struct {
	middleWares map[string]MiddleWare
}

func (router *Router)use(pattern string, mw MiddleWare) {

}
