package router

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2018-10-10

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPath2List(t *testing.T) {

	data := map[string][]string{
		"":                   nil,
		"/":                  []string{"/"},
		"//":                 []string{"/"},
		"///":                []string{"/"},
		"/test":              []string{"test"},
		"/test//":            []string{"test"},
		"/hello/world":       []string{"hello", "world"},
		"/hello/world/":      []string{"hello", "world"},
		"/hello/mr/:name":    []string{"hello", "mr", ":name"},
		"/hello/tail/*":      []string{"hello", "tail", "*"},
		"/hello/tail/*/":     []string{"hello", "tail", "*"},
		"/hello/tail/*//":    []string{"hello", "tail", "*"},
		"/hello/tail/*/test": nil,
	}

	for k, v := range data {
		res := path2list(k, true)

		if len(res) != len(v) || (res == nil && v != nil) || (res != nil && v == nil) {
			t.Fatal("failed " + k)
		}

		for i, item := range res {
			if item != v[i] {
				t.Fatal("failed " + k)
			}
		}
	}
}

var funcNum int = 0
var funcUid int64 = 0
var funcPid int64 = 0
var funcTail string = ""

func TestRouter(t *testing.T) {

	rules := map[string]HANDLER{
		"/users/:uid/posts": func(rw http.ResponseWriter, r *http.Request, p Params) {
			funcNum = 1
			funcUid = p.GetInt("uid")
			funcTail = ""
			rw.WriteHeader(200)
			rw.Write([]byte("1"))
		},

		"/tails/*": func(rw http.ResponseWriter, r *http.Request, p Params) {
			funcNum = 2
			funcUid = 0
			funcTail = p.GetString("*")
			rw.WriteHeader(200)
			rw.Write([]byte("1"))
		},

		"/users": func(rw http.ResponseWriter, r *http.Request, p Params) {
			funcNum = 3
			funcUid = 0
			funcTail = ""
			rw.WriteHeader(200)
			rw.Write([]byte("1"))
		},
		"/users/:uid/posts/:pid": func(rw http.ResponseWriter, r *http.Request, p Params) {
			funcNum = 4
			funcUid = p.GetInt("uid")
			funcPid = p.GetInt("pid")
			funcTail = ""
			rw.WriteHeader(200)
			rw.Write([]byte("1"))
		},
	}

	router := New(true)

	Register("DELETE", "/users/:uid/posts/:pid", func(rw http.ResponseWriter, r *http.Request, p Params) {
		funcNum = 5
		funcUid = p.GetInt("uid")
		funcPid = p.GetInt("pid")
		funcTail = ""
		rw.WriteHeader(200)
		rw.Write([]byte("1"))
	})

	Register("POST", "/users/:uid/posts", func(rw http.ResponseWriter, r *http.Request, p Params) {
		funcNum = 6
		funcUid = p.GetInt("uid")
		funcPid = 0
		funcTail = ""
		rw.WriteHeader(200)
		rw.Write([]byte("1"))
	})

	Register("POST", "/panic", func(rw http.ResponseWriter, r *http.Request, p Params) {
		funcNum = 7
		funcUid = 0
		funcPid = 0
		funcTail = ""
		panic("7")
	})

	Redirect("/redirect", "/dest", 302)

	router.OptionsFunc = func(rw http.ResponseWriter, req *http.Request) {
		funcNum = 8
		rw.WriteHeader(200)
	}

	for k, v := range rules {
		router.Register("GET", k, v)
	}

	req := httptest.NewRequest("GET", "/users/2/posts", nil)
	rw := httptest.NewRecorder()

	ServeHTTP(rw, req)
	if funcNum != 1 || rw.Result().StatusCode != 200 || funcUid != 2 || funcTail != "" {
		t.Fatal("/users/2/posts")
	}

	req = httptest.NewRequest("GET", "/tails/users/2/posts", nil)
	rw = httptest.NewRecorder()

	ServeHTTP(rw, req)
	if funcNum != 2 || rw.Result().StatusCode != 200 || funcUid != 0 || funcTail != "/users/2/posts" {
		t.Fatal("/tails/users/2/posts")
	}

	req = httptest.NewRequest("GET", "/tails", nil)
	rw = httptest.NewRecorder()

	ServeHTTP(rw, req)
	if rw.Result().StatusCode != 404 {
		t.Fatal("/tails")
	}

	req = httptest.NewRequest("GET", "/users/12/posts", nil)
	rw = httptest.NewRecorder()

	ServeHTTP(rw, req)
	if funcNum != 1 || rw.Result().StatusCode != 200 || funcUid != 12 || funcTail != "" {
		t.Fatal("/users/12/posts")
	}

	req = httptest.NewRequest("GET", "/userssssssssss", nil)
	rw = httptest.NewRecorder()

	ServeHTTP(rw, req)
	if rw.Result().StatusCode != 404 {
		t.Fatal("/usersssssssss")
	}

	req = httptest.NewRequest("GET", "/users", nil)
	rw = httptest.NewRecorder()

	ServeHTTP(rw, req)
	if rw.Result().StatusCode != 200 || funcNum != 3 || funcUid != 0 || funcTail != "" {
		t.Fatal("/users")
	}

	req = httptest.NewRequest("GET", "/users//", nil)
	rw = httptest.NewRecorder()

	ServeHTTP(rw, req)
	if rw.Result().StatusCode != 200 || funcNum != 3 || funcUid != 0 || funcTail != "" {
		t.Fatal("/users")
	}

	req = httptest.NewRequest("GET", "/users/122/posts/100", nil)
	rw = httptest.NewRecorder()

	ServeHTTP(rw, req)
	if funcNum != 4 || rw.Result().StatusCode != 200 || funcUid != 122 || funcPid != 100 || funcTail != "" {
		t.Fatal("/users/122/posts/100")
	}

	req = httptest.NewRequest("DELETE", "/users/122/posts/100", nil)
	rw = httptest.NewRecorder()

	ServeHTTP(rw, req)
	if funcNum != 5 || rw.Result().StatusCode != 200 || funcUid != 122 || funcPid != 100 || funcTail != "" {
		t.Fatal("/users/122/posts/100")
	}

	req = httptest.NewRequest("POST", "/users/22/posts/", nil)
	rw = httptest.NewRecorder()

	ServeHTTP(rw, req)
	if funcNum != 6 || rw.Result().StatusCode != 200 || funcUid != 22 || funcPid != 0 || funcTail != "" {
		t.Fatal("/users/22/posts/")
	}

	req = httptest.NewRequest("POST", "/panic", nil)
	rw = httptest.NewRecorder()

	ServeHTTP(rw, req)
	if rw.Result().StatusCode != 500 {
		t.Fatal("/panic")
	}

	req = httptest.NewRequest("OPTIONS", "/options", nil)
	rw = httptest.NewRecorder()

	ServeHTTP(rw, req)
	if rw.Result().StatusCode != 200 || funcNum != 8 {
		t.Fatal("/options")
	}

	req = httptest.NewRequest("GET", "/redirect", nil)
	rw = httptest.NewRecorder()

	ServeHTTP(rw, req)
	if rw.Result().StatusCode != 302 {
		t.Fatal("/redirect")
	}
}
