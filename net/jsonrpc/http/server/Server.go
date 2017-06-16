package server


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-06-16


import (
    "github.com/powerman/rpc-codec/jsonrpc2"
    "net"
    "net/http"
    "net/rpc"
    "strconv"
)


func Run( cfg *Config, classes ...interface{} ) {

  server := rpc.NewServer()

  for _, class := range classes {
    server.Register( class )
  }

  listener, err := net.Listen("tcp", cfg.Host + ":" + strconv.Itoa(cfg.Port) )

  if err != nil {
    panic(err)
  }

  defer listener.Close()

  http.Serve(listener, http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != cfg.Url {

      serverCodec := jsonrpc2.NewServerCodec(&HttpConn{in: r.Body, out: w}, server )
   
      w.Header().Set( "Content-type", "application/json" )
      w.WriteHeader(200)
   
      if err1 := server.ServeRequest(serverCodec) ; err1 != nil {
        http.Error(w, "Error while serving JSON request", 500)
        return
      }

    } else {
      http.Error(w, "Unknown request", 404)
    }

  } ) )
}

