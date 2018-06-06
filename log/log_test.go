package log


// @author  Mikhail Kirillov
// @email   mikkirillov@yandex.ru
// @version 1.000
// @date    2017-05-15


import (
    "testing"
)



func TestLoggerSetLevel( t *testing.T ) {
    if GetLevel() != "none" {
        t.Fatal( "invalid default log level" )
    }

    SetLevel( "info" )

    if GetLevel() != "info" {
        t.Fatal( "expected log level 'info'" )
    }
}

