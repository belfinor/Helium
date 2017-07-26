package ldb


// @author  Mikhail Kirillov
// @email   mikkirillov@yandex.ru
// @version 1.000
// @date    2017-05-18


import (
    "testing"
)


func TestDB( t *testing.T ) {
    Init( &Config{ Path: "/tmp/helium.test" } )
    
    key := []byte( "test.key" )
    val := []byte{ 1, 2, 3 }

    Del( key )

    res := Get( key )

    if res != nil {
        t.Fatal( "expecting nil value" )
    }

    Set( key, val )

    res = Get( key )

    if res == nil || len(res) != len(val) {
        t.Fatal( "db.Get error" )
    }

    for i, c := range res {
        if c != val[i] {
            t.Fatal( "db.Get error" )
        }
    }

    res = Get( key )

    if res == nil || len(res) != len(val) {
        t.Fatal( "db.Get error" )
    }

    for i, c := range res {
        if c != val[i] {
            t.Fatal( "db.Get error" )
        }
    }
}

