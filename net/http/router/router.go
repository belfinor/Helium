package router

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-10-02

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/belfinor/Helium/net/http/errors"
)

type Router struct {
	methods map[string]*node
}

func New() *Router {
	return &Router{
		methods: map[string]*node{},
	}
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

func (r *Router) Handler(rw http.ResponseWriter, req *http.Request) {

	root, has := r.methods[req.Method]
	if !has {
		errors.Send(rw, 404)
		return
	}

	paths := path2list(req.URL.Path, false)
	if paths == nil {
		errors.Send(rw, 400)
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

		errors.Send(rw, 404)
		return
	}

	if root.WildCard {
		params["*"] = tail
	}

	if root.F != nil {
		root.F(rw, req, Params(params))
	} else {
		errors.Send(rw, 404)
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
