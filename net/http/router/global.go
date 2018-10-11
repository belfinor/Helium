package router

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2018-10-11

import (
	"net/http"

	"github.com/belfinor/Helium/net/http/errors"
)

var handler *Router

func SetDefault(rt *Router) {
	handler = rt
}

func GetDefault() *Router {
	return handler
}

func Register(method string, path string, fn HANDLER) {
	if handler != nil {
		handler.Register(method, path, fn)
	}
}

func Redirect(from, to string, code int) {
	if handler != nil {
		handler.Redirect(from, to, code)
	}
}

func RegisterJsonRPC(url string) {
	if handler != nil {
		handler.RegisterJsonRPC(url)
	}
}

func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if handler != nil {
		handler.ServeHTTP(rw, req)
	} else {
		errors.Send(rw, 404)
	}
}
