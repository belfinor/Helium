package jsonrpc2


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-06


import (
    "bufio"
    "fmt"
    "github.com/belfinor/Helium/log"
    "net"
    "strconv"
    "strings"
)


var nextId int64 = 0;


func RunTcp( cfg *TcpConfig ) {

    ln, err := net.Listen( "tcp", cfg.Host + ":" + strconv.Itoa(cfg.Port) )
    
    if err != nil {
        panic( "bind port error" )
    }

    log.Info( "jsonrpc2 tcp server start " + cfg.Host + ":" + strconv.Itoa(cfg.Port) )

    for {
        
        conn, err := ln.Accept() 

        if err != nil {
            continue
        }
            
        nextId++

        go tcp_connet_handler(conn, nextId )
    }
}


func tcp_connet_handler(conn net.Conn, id int64 ) {
    defer conn.Close()

    log.Info( fmt.Sprintf( "jsonrpc2 new tcp conection #%d", id ) )

    reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(conn)

    for {
        
        str, err := reader.ReadString('\n')
        
        if err != nil {
            writer.Flush()
            log.Debug( fmt.Sprintf( "jsonrpc2 connection #%d broken", id ) )
            break
        }

        str = strings.TrimSpace(str)
        if str == "" {
            continue
        }

        resp := Handle( []byte(str) )

        if resp != nil {
            writer.WriteString( string(resp) + "\n" )
            writer.Flush()
        }
    }

    log.Info( fmt.Sprintf( "jsonrpc2 connection #%d closed", id ) )
}

