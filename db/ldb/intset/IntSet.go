package intset


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2017-07-26


import (
    "github.com/belfinor/Helium/db/ldb"
    "github.com/belfinor/Helium/pack"
)


var sync chan int = make( chan int, 1 )


func Create( key []byte, list ...int64 ) {
    if len(list) == 0 {
        ldb.Del( key )
    } else {
        ldb.Set( key, pack.IntList2Bytes(list) )
    }
}


func Remove( key []byte ) {
    ldb.Del( key )
}


func Push( key []byte, list ...int64 ) {

    sync <- 1

    items := pack.Bytes2IntList( ldb.Get( key ) )

    for _, val := range list {
        has := false
        
        for _, in := range items {
            if in == val {
                has = true
                break
            }
        }

        if !has {
            items = append( items, val )
        }
    }

    ldb.Set( key, pack.IntList2Bytes( items ) )

    <- sync
}


func Pop( key []byte, list ...int64 ) {

    sync <- 1

    items := pack.Bytes2IntList( ldb.Get( key ) )
    res   := make( []int64, 0, len(items) )

    for _, val := range items {
        has := false
        
        for _, in := range list {
            if in == val {
                has = true
                break
            }
        }

        if !has {
            res = append( res, val )
        }
    }

    ldb.Set( key, pack.IntList2Bytes( res ) )

    <- sync
}


func Get( key []byte ) []int64 {
    return pack.Bytes2IntList( ldb.Get( key ) )
}

