package ldb


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.008
// @date    2018-05-30


import (
    "github.com/syndtr/goleveldb/leveldb"
    "github.com/syndtr/goleveldb/leveldb/opt"
    "github.com/syndtr/goleveldb/leveldb/util"
    "github.com/belfinor/Helium/log"
)


type DB struct {
    ldb *leveldb.DB
}


type FOR_EACH_KEY_FUNC func([]byte) bool
type FOR_EACH_FUNC func([]byte,[]byte) bool


var _db *DB


var proxyConfig *Config


func Init( cfg *Config ) {
    if _db == nil {
        _db = &DB{
        }

        comp := opt.NoCompression
        mul  := 1
        if cfg.Compression {
            comp = opt.SnappyCompression
            mul = 2
        }

        size := cfg.FileSize * 1024 * 1024

        log.Info( "open database: " + cfg.Path )
        _db.ldb,_ = leveldb.OpenFile( cfg.Path, &opt.Options{
            CompactionTableSize: size,
            WriteBuffer:         size * mul,
            Compression:         comp,
            ReadOnly:            cfg.ReadOnly,
        } )
    }
}


func TestInit() {

  log.TestInit()

  cfg := &Config{ Path: "/var/tmp/ldb_test", Compression: true, FileSize: 16, ReadOnly: false }
  Init( cfg )

  for {
    list := List( []byte{}, 1000, 0, false )
    for _, key := range list {
      Del( key )
    }

    if len(list) < 1000 {
      break
    }
  }
}


func Close() {
  if _db != nil {
    _db.ldb.Close()
    _db = nil
  }
}


func InitProxy( cfg *Config ) {
  proxyConfig = cfg
}


func Use( base string ) {

  cfg := &Config{
    Path:        proxyConfig.Path + "/" + base,
    FileSize:    proxyConfig.FileSize,
    Compression: proxyConfig.Compression,
    ReadOnly:    proxyConfig.ReadOnly,
  }

  Close()
  Init( cfg )
}


func Source() *leveldb.DB {
  if _db == nil {
    return nil
  }

  return _db.ldb
}


func Set( key []byte, value []byte) {
    if value == nil || len(value) == 0 {
        _db.ldb.Delete( key, nil )
    } else {
        _db.ldb.Put( key, value, nil )
    }
}


func Get( key []byte ) []byte {
    val, err := _db.ldb.Get( key, nil )

    if err != nil {
        return nil
    }

    return val
}


func Has( key []byte ) bool {
    val, err := _db.ldb.Has( key, nil )
    if err != nil {
      return false
    }
    return val
}


func Del( key []byte ) {
    _db.ldb.Delete(key, nil)
}


func Total( prefix []byte) int64 {

    iter := _db.ldb.NewIterator(util.BytesPrefix(prefix), nil)
    defer iter.Release()

    i    := int64(0)

    for iter.Next() {
        i++
    }

    return i
}


func List( prefix []byte, limit int, offset int, RemovePrefix bool ) [][]byte {

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


func ForEach( prefix []byte, RemovePrefix bool, fn FOR_EACH_FUNC ) {

    iter := _db.ldb.NewIterator(util.BytesPrefix(prefix), nil)
    defer iter.Release()

    list := make( []byte, 4096 )
    var size int

    for iter.Next() {

        if( RemovePrefix ) {
            size = len( iter.Key() ) - len(prefix)
            copy( list, ( iter.Key() )[ len(prefix) : ] )
        } else {
            size = len( iter.Key() )
            copy( list, iter.Key() )
        }

        if !fn(list[:size], iter.Value() ) {
            return
        }
    }
}


func ForEachKey( prefix []byte, limit int, offset int, RemovePrefix bool, fn FOR_EACH_KEY_FUNC ) {

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

