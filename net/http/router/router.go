package router

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2018-10-03

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/net/http/errors"
	"github.com/belfinor/Helium/uniq"
)

func init() {
	New(true)
}

type Router struct {
	methods           map[string]*node
	NotFoundFunc      http.HandlerFunc
	BadRequestFunc    http.HandlerFunc
	InternalErrorFunc http.HandlerFunc
	OptionsFunc       http.HandlerFunc
	LoggerFunc        log.LoggerFunc
	MakeUid           bool
}

func New(isDefault bool) *Router {
	rt := &Router{
		methods:           map[string]*node{},
		NotFoundFunc:      func(rw http.ResponseWriter, req *http.Request) { errors.Send(rw, 404) },
		BadRequestFunc:    func(rw http.ResponseWriter, req *http.Request) { errors.Send(rw, 400) },
		InternalErrorFunc: func(rw http.ResponseWriter, req *http.Request) { errors.Send(rw, 500) },
		OptionsFunc:       nil,
		LoggerFunc:        nil,
		MakeUid:           true,
	}

	if isDefault {
		handler = rt
	}

	return rt
}

func (r *Router) Register(method string, path string, fn HANDLER) {

	root, has := r.methods[method]
	if !has {
		root = &node{Childs: make(map[string]*node)}
		r.methods[method] = root
	}

	list := path2list(path, true)
	if list == nil {
		return
	}

	for _, item := range list {

		if item[0] == ':' {
			n, h := root.Childs[""]
			if !h {
				name := item
				if len(name) > 1 {
					name = item[1:]
				} else {
					name = ""
				}
				n = &node{Name: name, WildCard: false, Childs: make(map[string]*node)}
			}
			root.Childs[""] = n
			root = n
			continue
		}

		if item == "*" {
			_, h := root.Childs[""]
			if !h {
				root.Childs[""] = &node{Name: "*", F: fn, WildCard: true}
			}
			return
		}

		n, h := root.Childs[item]
		if !h {
			n = &node{Name: "", WildCard: false, Childs: make(map[string]*node)}
		}
		root.Childs[item] = n
		root = n

	}

	root.F = fn
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	if r.MakeUid {
		makeUid(rw, req)
	}

	if r.LoggerFunc != nil {
		writeLog(r.LoggerFunc, req.Method, req.RequestURI)
	}

	defer func() {

		if re := recover(); re != nil {
			log.Error(re)
			r.InternalErrorFunc(rw, req)
		}

	}()

	if req.Method == http.MethodOptions && r.OptionsFunc != nil {
		r.OptionsFunc(rw, req)
		return
	}

	root, has := r.methods[req.Method]
	if !has {
		r.NotFoundFunc(rw, req)
		return
	}

	paths := path2list(req.URL.Path, false)
	if paths == nil {
		r.BadRequestFunc(rw, req)
		return
	}

	tail := ""
	params := make(map[string]string)

	for _, item := range paths {

		if root.WildCard {
			tail += "/" + item
			continue
		}

		if n, h := root.Childs[""]; h {
			root = n
			if root.WildCard {
				tail = "/" + item
			} else {
				params[root.Name] = item
			}
			continue
		}

		if n, h := root.Childs[item]; h {
			root = n
			continue
		}

		r.NotFoundFunc(rw, req)
		return
	}

	if root.WildCard {
		params["*"] = tail
	}

	if root.F != nil {
		root.F(rw, req, Params(params))
	} else {
		r.NotFoundFunc(rw, req)
	}
}

func path2list(path string, regmod bool) []string {
	rh := strings.NewReader(path)
	builder := strings.Builder{}
	mode := 0
	result := make([]string, 0, 8)

	for {
		run, _, err := rh.ReadRune()
		if err != nil {
			break
		}

		switch mode {
		case 0:

			if run != '/' {
				return nil
			}

			mode = 1

		case 1:

			if run != '/' {
				builder.WriteRune(run)
				mode = 2
			}

		case 2:

			if run == '/' {
				str := builder.String()
				builder.Reset()

				if str == "*" {
					result = append(result, "*")
					if regmod {
						mode = 3
					} else {
						mode = 1
					}
				} else {

					uri, e := url.PathUnescape(str)
					if e != nil {
						return nil
					}

					result = append(result, uri)
					mode = 1
				}
			} else {
				builder.WriteRune(run)
			}

		case 3:

			if run != '/' {
				return nil
			}
		}
	}

	switch mode {

	case 0:
		return nil

	case 1:
		if len(result) == 0 {
			result = append(result, "/")
		}

	case 2:
		str := builder.String()
		builder.Reset()

		uri, e := url.PathUnescape(str)
		if e != nil {
			return nil
		}

		result = append(result, uri)
	}

	return result
}

func writeLog(f log.LoggerFunc, method string, url string) {
	f(fmt.Sprintf("%s %s", method, url))
}

var UNIQS *uniq.Uniq = uniq.New()

func makeUid(rw http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("uid")

	v := ""

	if err != nil {
		v = UNIQS.Next()
	} else {
		v = c.Value
	}

	cookie := &http.Cookie{
		Name:     "uid",
		Value:    v,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(time.Now().Unix()+86400*366, 0),
	}

	req.AddCookie(cookie)

	http.SetCookie(rw, cookie)
}
