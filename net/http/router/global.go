package router

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-10-02

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

func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if handler != nil {
		handler.ServeHTTP(rw, req)
	} else {
		errors.Send(rw, 404)
	}
}
