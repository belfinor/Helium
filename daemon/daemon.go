package daemon

// @author  Mikhail Kirillov
// @email   mikkirillov@yandex.ru
// @version 1.000
// @date    2017-05-24

import (
	"fmt"
	"github.com/sevlyar/go-daemon"
	"os"
	"time"
)

func Run(conf *Config) {

	var cntxt *daemon.Context

	cntxt = &daemon.Context{
		PidFileName: conf.PidFile,
		PidFilePerm: 0644,
		LogFileName: conf.LogFile,
		LogFilePerm: 0640,
		WorkDir:     conf.WordDir,
		Umask:       027,
		Args:        os.Args,
	}

	child, err := cntxt.Reborn()

	if err != nil {
		fmt.Println(err)
	}

	if child != nil {
		time.Sleep(time.Second)
		os.Exit(0)
	}
}
