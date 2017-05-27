package counter


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-05-27



import (
    "testing"
    "github.com/belfinor/Helium/test"
)



func TestCounter( t *testing.T ) {
    test.Init()

    key := []byte( "test.cnt" )

    Reset( key )
    if Get(key) != 0 {
        t.Fatal( "Reset not work" )
    }
    
    if Inc(key,2) != 2 {
        t.Fatal( "Inc not work" )
    }
    
    if Inc(key,3) != 5 {
        t.Fatal( "Inc not work" )
    }

    Set( key, int64(12) )
    if Get(key) != 12 {
        t.Fatal( "Set not work" )
    } 
}


