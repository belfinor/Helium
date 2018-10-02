package wdog

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-10-02

import (
	"os"
	"time"

	"github.com/belfinor/Helium/log"
)

type WatchDog struct {
	input chan bool
	done  chan bool
	tm    time.Duration
}

func New(keepalive time.Duration) *WatchDog {
	wd := &WatchDog{
		input: make(chan bool, 10),
		done:  make(chan bool),
		tm:    keepalive,
	}

	go func() {
		for {

			select {

			case _, ok := <-wd.input:

				if !ok {
					close(wd.done)
					return
				}

			case <-time.After(wd.tm):

				log.Error("wdog timeout. terminate application")
				<-time.After(time.Second * 2)
				os.Exit(1)
			}

		}
	}()

	return wd
}

func (wd *WatchDog) Alive() {
	wd.input <- true
}

func (wd *WatchDog) Close() {
	close(wd.input)
	<-wd.done
}
