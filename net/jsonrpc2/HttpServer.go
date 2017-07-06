package jsonrpc2


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-06


import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strconv"
)


func httpHandler( rw http.ResponseWriter, req *http.Request ) {

    if req.Method != "POST" {
        rw.WriteHeader(http.StatusNotFound)
        fmt.Fprint( rw, "404 page not found" )
        return
    }

    body, _ := ioutil.ReadAll(req.Body)

    resp := Handle( body )

    if resp != nil {
        rw.Header().Set( "Content-Type", "application/json; charset=utf-8" )
        fmt.Fprint( rw, string(resp) )
    }
}


func RunHttp( cfg *HttpConfig ) {
    http.HandleFunc( cfg.Url, httpHandler )
    http.ListenAndServe( cfg.Host + ":" + strconv.Itoa(cfg.Port), nil)
}

