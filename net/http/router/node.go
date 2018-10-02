package router

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-10-02

import (
	"net/http"
)

type HANDLER func(rw http.ResponseWriter, r *http.Request, params Params)

type node struct {
	Name     string
	Childs   map[string]*node
	WildCard bool
	F        HANDLER
}
