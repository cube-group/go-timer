package loader

import (
	"net/url"
)

const (
	LOAD_TYPE_MULTIPLE = 0
	LOAD_TYPE_SINGLE = 1

	METHOD_GET = "get"
	METHOD_POST = "post"
)

type URLRequest struct {
	Url     string
	Method  string
	Headers URLHeaders
	Body    URLVariables
}

type URLHeaders struct {
	url.Values
}

type URLEvent struct {
	content string
	status  string
	error   string
}

type URLHandler func(e URLEvent)

type URLVariables map[string]string

type URLLoader struct {
	Request URLRequest
	Handler URLHandler
	Type    int //details in LOAD_TYPE_*
}

func (loader *URLLoader)start() {
	go startLoad(loader)
}

func startLoad(loader *URLLoader) {
	//coming soon..
}

func load(url string, method string, body interface{}, handler URLHandler) {

	var vars URLVariables

	if (method == METHOD_POST) {
		switch body.(type) {
		case map[string]string:
			vars = body.(URLVariables)
		}
	}

	request := URLRequest{
		Url:url,
		Method:method,
		Headers:nil,
		Body:vars,
	}

	loader := URLLoader{
		Request:request,
		Handler:handler,
		Type:LOAD_TYPE_MULTIPLE,
	}

	loader.start()
}