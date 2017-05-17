package pack


// @author  Mikhail Kirillov
// @version 1.000
// @date    2017-05-17


import (
    "testing"
)


func TestInt2Bytes( t *testing.T ) {
    val := int64(1234567890)
    b := Int2Bytes( val )
    if val != Bytes2Int(b) {
        t.Fatal(Bytes2Int(b))
        t.Fatal( "expecting 1234567890" )
    }
    
}


func TestIntList2Bytes( t *testing.T ) {
    list := []int64{ 1, 129, 1045, 76532, 462784628 }
    b := IntList2Bytes( list )
    res := Bytes2IntList(b)

    if len(res) != len(list) {
        t.Fatal( "ivalid decoded size" )
    }

    for i, cur := range res {
        if cur != list[i] {
            t.Fatalf( "expected %d", cur )
        }
    }
}


