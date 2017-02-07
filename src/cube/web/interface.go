package web

import "time"

type ICookie interface {
	get(string) string
	set(string, string, time.Duration)
	delete(string)
	clear()
}

type ISession interface {
	sessionID(string)
	sessionName(string)
	get(string) string
	set(string, string, time.Duration)
	delete(string)
	clear()
}

type IBody interface {
	get(string) string
}

type IQuery interface {
	get(string) string
}

type IRequest interface {
	header(string) string
}

type IResponse interface {
	display(templateName string, interface{}) error
	render(templateName string, interface{}) byte
	download(filePath string) error
	send(string) error
	json(string) error
	jsonp(string) error

	status(int) IResponse
	header(string, string) IResponse
}

type IRouter interface {
	use(string, MiddleWare)
}
