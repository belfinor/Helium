package server


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-05-24


import (
    "bufio"
    "fmt"
    "github.com/belfinor/Helium/log"
    "net"
    "strconv"
)


type CALLBACK_FUNC func(string) ([]string,bool)


type Server struct {
    Port         int
    Host         string
    Callback     CALLBACK_FUNC
}


var prefix string = "net/tcp/server: ";
var nextId int64 = 1;


func (s *Server) Start() {

	ln, err := net.Listen( "tcp", s.Host + ":" + strconv.Itoa(s.Port) )
	
	if err != nil {
		panic( "bind port error" )
	}

        log.Info( prefix + "start " + s.Host + ":" + strconv.Itoa(s.Port) )

	for {
		conn, err := ln.Accept()
		
		if err != nil {
			continue
		}
		
                nextId++
		go s.connet_handler(conn, nextId )
	}
}


func (s *Server) connet_handler(conn net.Conn, id int64 ) {
    defer conn.Close()

    log.Info( prefix + fmt.Sprintf( "new conection #%d", id ) )

    reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(conn)

    for {
        
        str, err := reader.ReadString('\n')
        
        if err != nil {
            writer.Flush()
            log.Debug( prefix + fmt.Sprintf( "connection #%d broken", id ) )
            break
        }

        resp, stop := s.Callback( str )

        if resp != nil {
            if len(resp) == 0 {
                resp = []string{ "(no data)" }
            }

            for i, val := range resp {
                if i == len(resp) - 1 {
                    writer.WriteString( val + "\n" )
                } else {
                    writer.WriteString( "<<< " + val + "\n" )
                }
                writer.Flush()
            }
        }

        if stop {
            break
        }
    }

    log.Debug( prefix + fmt.Sprintf( "connection #%d closed" ) )
}

