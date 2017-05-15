package log


// @author  Mikhail Kirillov
// @email   mikkirillov@yandex.ru
// @version 1.000
// @date    2017-05-15


import (
    l "log"
)


var _log_level int = 0


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
        l.Println( level + ": " + text )
    }
}


func Fatal( str string ) {
    logger( "fatal", str )
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

