package wdog

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-10-01

import (
	"os"
	"time"

	"github.com/belfinor/Helium/log"
)

var input chan bool = make(chan bool)

func Run(timeout time.Duration) {

	for {

		select {
		case <-input:

		case <-time.After(timeout):

			log.Error("wdog tmeout. terminate application")
			<-time.After(time.Second * 2)
			os.Exit(1)
		}

	}

}

func Alive() {
	input <- true
}
