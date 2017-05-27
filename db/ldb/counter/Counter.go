package counter


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-05-27


import (
    "github.com/belfinor/Helium/db/ldb"
    "github.com/belfinor/Helium/pack"
)


func Inc( key []byte, val int64 ) int64 {
    ldb.Lock()
    defer ldb.Unlock()

    raw := ldb.GetUnsafe( key )
    cur := pack.Bytes2Int( raw )

    cur += val

    if cur <= int64(0) {
        ldb.DelUnsafe( key )
        cur = 0
    } else {
        ldb.SetUnsafe( key, pack.Int2Bytes(cur) )
    }

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

