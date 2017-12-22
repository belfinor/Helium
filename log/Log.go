package log


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2017-05-29


import (
    "github.com/belfinor/Helium/time/strftime"
    "os"
    "time"
)


var _log_level int = 0


type Config struct {
    Template string `json:"template"`
    Period   int    `json:"period"`
    Save     int    `json:"save"`
    Level    string `json:"level"`
}



var _config *Config
var input chan string
var fh *os.File
var filename string
var lastCheck int64


var _log_levels map[string]int = map[string]int{
    "none":  0,
    "fatal": 1,
    "error": 2,
    "info":  3,
    "debug": 4,
    "trace": 5,
}


func logger( level string, text string ) {
    code, ok := _log_levels[level]
    if ok && code <= _log_level {
        input <-  level + "| " + text
    }
}


func Init( c *Config ) {
    if _config == nil {

        _config = c
        filename = strftime.Format( c.Template, time.Now() )
        var err error

        if fh, err = os.OpenFile(filename, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0755) ; err != nil {
            panic(err)
        }

        input = make( chan string, 1024 )
        lastCheck = time.Now().Unix()
 
        SetLevel( c.Level )
       
        if _config.Save > 0 {
            rm_name := strftime.Format( _config.Template, time.Unix( lastCheck - int64(_config.Save * _config.Period), 0 ) )
            os.Remove(rm_name)
        }

        go logWriter()
    }
}


func logRotate() {

    if lastCheck + 60 > time.Now().Unix() {
        return
    }

    lastCheck = time.Now().Unix()
    new_name := strftime.Format( _config.Template, time.Now() )

    if new_name != filename {
        fh.Close()
                
        var err error
        filename = new_name       
         
        if fh, err = os.OpenFile(filename, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0755) ; err != nil {
            panic(err)
        }
                
        if _config.Save > 0 {
            rm_name := strftime.Format( _config.Template, time.Unix( lastCheck - int64(_config.Save * _config.Period), 0 ) )
            os.Remove(rm_name)
        }
    }
}


func logWriter() {
    for {
        select {
        
        case str := <- input:
            logRotate()    
   
            fh.WriteString( strftime.Format( "%Y-%m-%d %H:%M:%S", time.Now() ) + "|" + str + "\n" )
            fh.Sync()
 
        case <- time.After( time.Minute ):
            logRotate()    
        }
    }
}



func Fatal( str string ) {
    logger( "fatal", str )
    <- time.After( time.Second * 2 )
    os.Exit(1)
}


func Error( str string ) {
    logger( "error", str )
}


func Info( str string ) {
    logger( "info", str )
}


func Debug( str string ) {
    logger( "debug", str )
}


func Trace( str string ) {
    logger( "trace", str )
}


func SetLevel( level string ) {
    code, ok := _log_levels[level]

    if ok {
        _log_level = code
    } else {
        _log_level = 0
    }
}


func GetLevel() string {
    for code, level := range _log_levels {
        if level == _log_level {
            return code
        }
    }
    return "none"
}

