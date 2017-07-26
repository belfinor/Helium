package ldb


// @author  Mikhail Kirillov
// @email   mikkirillov@yandex.ru
// @version 1.001
// @date    2017-05-26


import (
    "github.com/syndtr/goleveldb/leveldb"
    "github.com/syndtr/goleveldb/leveldb/opt"
    "github.com/syndtr/goleveldb/leveldb/util"
    "github.com/belfinor/Helium/log"
    "sync"
)


type DB struct {
    sync.Mutex
    dbpath string
    ldb *leveldb.DB
}


type FOR_EACH_KEY_FUNC func([]byte) bool


var _db *DB


// must set before Init
func EnableSnappy( f bool ) {
    if f {
      opt.DefaultCompressionType = opt.SnappyCompression
    } else {
      opt.DefaultCompressionType = opt.NoCompression
    }
}


// must set before Init
func SetTableSize( mb int ) {
    opt.DefaultCompactionTableSize = mb * opt.MiB
}


func Init( base string ) {
    if _db == nil {
        _db = &DB{
            dbpath: base,
        }

        log.Info( "open database: " + base )
        _db.ldb,_ = leveldb.OpenFile(base,nil)   
    }
}


func SetUnsafe( key []byte, value []byte) {
    if value == nil || len(value) == 0 {
        _db.ldb.Delete( key, nil )
    } else {
        _db.ldb.Put( key, value, nil )
    }
}


func Set( key []byte, value []byte ) {
    _db.Lock()
    defer _db.Unlock()

    SetUnsafe( key, value )
}


func GetUnsafe( key []byte ) []byte {
    val, err := _db.ldb.Get( key, nil )

    if err != nil {
        return nil
    }

    return val
}


func Get( key []byte ) []byte {
    _db.Lock()
    defer _db.Unlock()

    return GetUnsafe( key )
}


func Del( key []byte ) {
    _db.Lock()
    defer _db.Unlock()
    DelUnsafe( key )
}


func DelUnsafe( key []byte ) {
    _db.ldb.Delete(key, nil)
}


func Lock() {
    _db.Lock()
}


func Unlock() {
    _db.Unlock()
}


func TotalUnsafe( prefix []byte) int64 {

    iter := _db.ldb.NewIterator(util.BytesPrefix(prefix), nil)
    defer iter.Release()

    i    := int64(0)

    for iter.Next() {
        i++
    }

    return i
}


func Total( prefix []byte ) int64 {
    _db.Lock()
    defer _db.Unlock()
    return TotalUnsafe( prefix )
}


func ListUnsafe( prefix []byte, limit int, offset int, RemovePrefix bool ) [][]byte {

    iter := _db.ldb.NewIterator(util.BytesPrefix(prefix), nil)
    defer iter.Release()

    res  := make( [][]byte, 0 )
    i    := -1

    for iter.Next() {
        i++

        if i >= offset + limit {
            break
        }

        if i < offset {
            continue
        }

        var list []byte

        if( RemovePrefix ) {
            size := len( iter.Key() ) - len(prefix)
            list = make( []byte, size )
            copy( list, ( iter.Key() )[ len(prefix) : ] )
        } else {
            list = make( []byte, len( iter.Key() ) )
            copy( list, iter.Key() )
        }

        res = append( res, list )
    }

    return res
}


func List( prefix []byte, limit int, offset int, RemovePrefix bool ) [][]byte {
    _db.Lock()
    defer _db.Unlock()
    return ListUnsafe( prefix, limit, offset, RemovePrefix )
}


func ForEachKey( prefix []byte, limit int, offset int, RemovePrefix bool, fn FOR_EACH_KEY_FUNC ) {
    _db.Lock()
    defer _db.Unlock()
    ForEachKeyUnsafe( prefix, limit, offset, RemovePrefix, fn )
}


func ForEachKeyUnsafe( prefix []byte, limit int, offset int, RemovePrefix bool, fn FOR_EACH_KEY_FUNC ) {

    iter := _db.ldb.NewIterator(util.BytesPrefix(prefix), nil)
    defer iter.Release()

    i    := -1

    for iter.Next() {
        i++

        if i >= offset + limit {
            break
        }

        if i < offset {
            continue
        }

        var list []byte

        if( RemovePrefix ) {
            size := len( iter.Key() ) - len(prefix)
            list = make( []byte, size )
            copy( list, ( iter.Key() )[ len(prefix) : ] )
        } else {
            list = make( []byte, len( iter.Key() ) )
            copy( list, iter.Key() )
        }

        if !fn(list) {
            return
        }
    }
}

