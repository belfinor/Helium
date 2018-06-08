package log


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.008
// @date    2018-06-08


import (
    "fmt"
    "github.com/belfinor/Helium/time/strftime"
    "os"
    "time"
)


var logLevel int = 0


type Config struct {
    Template string `json:"template"`
    Period   int    `json:"period"`
    Save     int    `json:"save"`
    Level    string `json:"level"`
    StdOut   bool   `json:"stdout"`
    StdErr   bool   `json:"stderr"`
}



var conf *Config
var input chan string = make( chan string, 1024 )
var fh *os.File
var filename string
var lastCheck int64 = time.Now().Unix()
var eofc chan bool = make( chan bool )


var logLevels map[string]int = map[string]int{
    "none":  0,
    "fatal": 1,
    "error": 2,
    "warn":  3,
    "info":  4,
    "debug": 5,
    "trace": 6,
}


func logger( level string, strs []interface{} ) {
    code, ok := logLevels[level]
    if ok && code <= logLevel {
        for _, text := range strs {
          input <-  level + "| " + fmt.Sprint(text)
        }
    }
}


func Init( c *Config ) {
    if conf == nil {

        conf = c
        filename = strftime.Format( c.Template, time.Now() )
        var err error

        if fh, err = os.OpenFile(filename, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0755) ; err != nil {
            panic(err)
        }

        SetLevel( c.Level )

        if conf.Save > 0 {
            rm_name := strftime.Format( conf.Template, time.Unix( lastCheck - int64(conf.Save * conf.Period), 0 ) )
            os.Remove(rm_name)
        }

        go logWriter()
    }
}


func TestInit() {
  if conf == nil {
    conf = &Config{ Template: "test.log", Period: 86400, Save: 20, Level: "none" }
    SetLevel( "none" )
    go logWriter()
  }
}


func logRotate() {

    if lastCheck + 60 > time.Now().Unix() {
        return
    }

    lastCheck = time.Now().Unix()
    new_name := strftime.Format( conf.Template, time.Now() )

    if new_name != filename {
        fh.Close()

        var err error
        filename = new_name

        if fh, err = os.OpenFile(filename, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0755) ; err != nil {
            panic(err)
        }

        if conf.Save > 0 {
            rm_name := strftime.Format( conf.Template, time.Unix( lastCheck - int64(conf.Save * conf.Period), 0 ) )
            os.Remove(rm_name)
        }
    }
}


func logWriter() {
    for {
        select {

        case str := <- input:

            if str == "eof" {
              eofc <- true
              return
            }

            logRotate()

            str = strftime.Format( "%Y-%m-%d %H:%M:%S", time.Now() ) + "|" + str + "\n"
            fh.WriteString( str )
            fh.Sync()

            if conf.StdOut {
              os.Stdout.WriteString( str )
              os.Stdout.Sync()
            }

            if conf.StdErr {
              os.Stderr.WriteString( str )
              os.Stderr.Sync()
            }

        case <- time.After( time.Minute ):
            logRotate()
        }
    }
}



func Fatal( str ...interface{} ) {
    logger( "fatal", str )
    input <- "eof"
    <- eofc
    os.Exit(1)
}


func Finish( str ...interface{} ) {
    logger( "info", str )
    input <- "eof"
    <- eofc
}


func Error( str ...interface{} ) {
    logger( "error", str )
}


func Info( str ...interface{} ) {
    logger( "info", str )
}


func Debug( str ...interface{} ) {
    logger( "debug", str )
}


func Warn( str ...interface{} ) {
    logger( "warn", str )
}


func Trace( str ...interface{} ) {
    logger( "trace", str )
}


func SetLevel( level string ) {
    code, ok := logLevels[level]

    if ok {
        logLevel = code
    } else {
        logLevel = 0
    }
}


func GetLevel() string {
    for code, level := range logLevels {
        if level == logLevel {
            return code
        }
    }
    return "none"
}

