package log

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.010
// @date    2018-10-02

import (
	"fmt"
	"os"
	"time"

	"github.com/belfinor/Helium/time/strftime"
)

var logLevels map[string]int = map[string]int{
	"none":  0,
	"fatal": 1,
	"error": 2,
	"warn":  3,
	"info":  4,
	"debug": 5,
	"trace": 6,
}

type LoggerFunc func(str ...interface{})

type Log struct {
	level     int
	conf      *Config
	input     chan string
	fh        *os.File
	filename  string
	lastCheck int64
	eofC      chan bool
}

func (l *Log) logger(level string, strs []interface{}) {

	if l == nil {
		return
	}

	code, ok := logLevels[level]
	if ok && code <= l.level {
		for _, text := range strs {
			l.input <- level + "| " + fmt.Sprint(text)
		}
	}
}

func New(c *Config, def bool) (*Log, error) {

	if def && defLog != nil {
		return defLog, nil
	}

	l := &Log{
		conf:      c,
		filename:  strftime.Format(c.Template, time.Now()),
		lastCheck: time.Now().Unix(),
		input:     make(chan string, 1024),
		eofC:      make(chan bool),
	}

	var err error

	if l.fh, err = os.OpenFile(l.filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755); err != nil {
		return nil, err
	}

	l.SetLevel(c.Level)

	if c.Save > 0 {
		rm_name := strftime.Format(c.Template, time.Unix(l.lastCheck-int64(c.Save*c.Period), 0))
		os.Remove(rm_name)
	}

	if def && defLog == nil {
		defLog = l
	}

	go l.writer()

	return l, nil
}

func Init(c *Config) {
	New(c, true)
}

func (l *Log) rotate() {

	if l.lastCheck+60 > time.Now().Unix() {
		return
	}

	l.lastCheck = time.Now().Unix()
	new_name := strftime.Format(l.conf.Template, time.Now())

	if new_name != l.filename {
		l.fh.Close()

		var err error
		l.filename = new_name

		if l.fh, err = os.OpenFile(l.filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755); err != nil {
			panic(err)
		}

		if l.conf.Save > 0 {
			rm_name := strftime.Format(l.conf.Template, time.Unix(l.lastCheck-int64(l.conf.Save*l.conf.Period), 0))
			os.Remove(rm_name)
		}
	}
}

func (l *Log) writer() {
	for {
		select {

		case str := <-l.input:

			if str == "eof" {
				l.eofC <- true
				return
			}

			l.rotate()

			str = strftime.Format("%Y-%m-%d %H:%M:%S", time.Now()) + "|" + str + "\n"
			l.fh.WriteString(str)
			l.fh.Sync()

			conf := l.conf

			if conf.StdOut {
				os.Stdout.WriteString(str)
				os.Stdout.Sync()
			}

			if conf.StdErr {
				os.Stderr.WriteString(str)
				os.Stderr.Sync()
			}

		case <-time.After(time.Minute):
			l.rotate()
		}
	}
}

func (l *Log) Close() {
	if l != nil {
		l.input <- "eof"
		<-l.eofC
	}
}

func (l *Log) Fatal(str ...interface{}) {
	l.logger("fatal", str)
	l.Close()
	os.Exit(1)
}

func (l *Log) Finish(str ...interface{}) {
	l.logger("info", str)
	l.Close()
}

func (l *Log) Error(str ...interface{}) {
	l.logger("error", str)
}

func (l *Log) Info(str ...interface{}) {
	l.logger("info", str)
}

func (l *Log) Debug(str ...interface{}) {
	l.logger("debug", str)
}

func (l *Log) Warn(str ...interface{}) {
	l.logger("warn", str)
}

func (l *Log) Trace(str ...interface{}) {
	l.logger("trace", str)
}

func (l *Log) SetLevel(lvl string) {
	if l == nil {
		return
	}

	if code, ok := logLevels[lvl]; ok {
		l.level = code
	} else {
		l.level = 0
	}
}

func (l *Log) GetLevel() string {
	if l == nil {
		return "none"
	}

	for code, level := range logLevels {
		if level == l.level {
			return code
		}
	}
	return "none"
}
