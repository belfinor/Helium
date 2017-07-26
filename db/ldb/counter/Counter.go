package counter


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2017-07-26


import (
    "github.com/belfinor/Helium/db/ldb"
    "github.com/belfinor/Helium/pack"
)


var sync chan int = make( chan int, 1 )


func Inc( key []byte, val int64 ) int64 {

    sync <- 1

    raw := ldb.Get( key )
    cur := pack.Bytes2Int( raw )

    cur += val

    if cur <= int64(0) {
        ldb.Del( key )
        cur = 0
    } else {
        ldb.Set( key, pack.Int2Bytes(cur) )
    }

    <- sync

    return cur
}


func Get( key []byte ) int64 {
    raw := ldb.Get( key )
    cur := pack.Bytes2Int(raw)
    return cur
}


func Reset( key []byte ) {
    ldb.Del( key )
}


func Set( key []byte, val int64 ) {
    if val <= int64(0) {
        ldb.Del( key )
    } else {
        ldb.Set( key, pack.Int2Bytes(val) )
    }
}

