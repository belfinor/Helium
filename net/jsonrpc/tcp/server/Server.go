package server


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-06-16


import (
  "context"
  "net"
  "net/rpc"
  "github.com/powerman/rpc-codec/jsonrpc2"
  "strconv"
)


var contextKey = "RemoteAddr"


func Run( cfg *Config, classes ...interface{} ) {
    
  for _, class := range classes {
    rpc.Register( class )
  }

  lnTCP, err := net.Listen("tcp", cfg.Host + ":" + strconv.Itoa(cfg.Port) )
    
  if err != nil {
    panic(err)
  }
    
  defer lnTCP.Close()
    
  for {
    conn, err := lnTCP.Accept()
    if err != nil {
      return
    }
    
    ctx := context.WithValue(context.Background(), contextKey, conn.RemoteAddr())
    go jsonrpc2.ServeConnContext(ctx, conn)  
  }    
}

