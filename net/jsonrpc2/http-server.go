package jsonrpc2

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-10-02

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/belfinor/Helium/log"
)

func HttpHandler(rw http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprint(rw, "404 page not found")
		return
	}

	body, _ := ioutil.ReadAll(req.Body)

	resp := Handle(body)

	if resp != nil {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprint(rw, string(resp))
	}
}

func RunHttp(cfg *HttpConfig) {
	http.HandleFunc(cfg.Url, HttpHandler)
	log.Info("start jsonrpc2 server addr=" + cfg.Host + ":" + strconv.Itoa(cfg.Port) + " url=" + cfg.Url)
	http.ListenAndServe(cfg.Host+":"+strconv.Itoa(cfg.Port), nil)
}
