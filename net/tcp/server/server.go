package server

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2017-05-29

import (
	"bufio"
	"fmt"
	"github.com/belfinor/Helium/log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type Context struct {
	Id     int64
	State  map[string]string
	Random *rand.Rand
	Tact   int64
}

type CALLBACK_FUNC func(string, *Context) ([]string, bool)

type Server struct {
	Port     int
	Host     string
	Callback CALLBACK_FUNC
}

var prefix string = "net/tcp/server: "
var nextId int64 = 0

func (s *Server) Start() {

	ln, err := net.Listen("tcp", s.Host+":"+strconv.Itoa(s.Port))

	if err != nil {
		panic("bind port error")
	}

	log.Info(prefix + "start " + s.Host + ":" + strconv.Itoa(s.Port))

	for {

		conn, err := ln.Accept()

		if err != nil {
			continue
		}

		nextId++

		go s.connet_handler(conn, nextId)
	}
}

func (s *Server) connet_handler(conn net.Conn, id int64) {
	defer conn.Close()

	context := Context{
		Id:     id,
		State:  make(map[string]string),
		Random: rand.New(rand.NewSource(time.Now().Unix())),
		Tact:   0,
	}

	log.Info(prefix + fmt.Sprintf("new conection #%d", id))

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {

		str, err := reader.ReadString('\n')

		if err != nil {
			writer.Flush()
			log.Debug(prefix + fmt.Sprintf("connection #%d broken", id))
			break
		}

		context.Tact++

		resp, stop := s.Callback(str, &context)

		if resp != nil {
			if len(resp) == 0 {
				resp = []string{"(no data)"}
			}

			for i, val := range resp {
				if i == len(resp)-1 {
					writer.WriteString(val + "\n")
				} else {
					writer.WriteString("<<< " + val + "\n")
				}
				writer.Flush()
			}
		}

		if stop {
			log.Debug(prefix + fmt.Sprintf("conection #%d get stop signal", id))
			break
		}
	}

	log.Info(prefix + fmt.Sprintf("connection #%d closed", id))
}
