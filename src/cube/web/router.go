package web

import (
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

	r        *http.Request
}

//get header from the Request
func (req Request)header(key string) string {
	return req.r.Header.Get(key)
}

//package the new response
type Response struct {
	writer http.ResponseWriter
}

type Router struct {
	pattern string
	req     Request
	res     Response
}

func NewRouter(writer http.ResponseWriter, r *http.Request) Router {
	req := Request{
		Protocol:r.Proto,
		Host:r.Host,
		Refer:r.Referer(),
		Method:r.Method,

		r:r,
	}

	res := Response{
		writer:writer,
	}

	router := Router{r.URL.Path, req, res}

	return router
}
