package web

import (
	"container/list"
	"net/http"
)


//package the new request
type Request struct {
	Protocol string
	Host     string
	Refer    string
	Uri      string

	Method   string

	Cookie   ICookie
	Session  ISession

	Body     IBody
	Query    IQuery
}

//get header from the Request
func (req Request)header(key string) string {
	return req.dispatcher.r.Header.Get(key)
}

//package the new response
type Response struct {
}

type Router struct {
	pattern string
	req     Request
	res     Response
}

func NewRouter(r *http.Request) Router {
	req := Request{
		Protocol:r.Proto,
		Host:r.Host,
		Refer:r.Referer(),
		Method:r.Method,
	}

	res := Response{
	}

	router := Router{r.URL.Path, req, res}

	return router
}
