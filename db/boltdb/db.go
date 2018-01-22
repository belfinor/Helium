package boltdb


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-01-22


import (
  "github.com/belfinor/Helium/log"
  "github.com/boltdb/bolt"
  "os"
  "time"
)


type DB struct {
  db *bolt.DB
}


func Open( cfg *Config ) ( *DB, error ) {

  opts := &bolt.Options{ ReadOnly: cfg.ReadOnly }

  if cfg.Timeout != 0 {
    opts.Timeout = time.Duration( int64(cfg.Timeout) * int64(time.Second) )
  }

  db, err := bolt.Open( cfg.Database, os.FileMode(cfg.Mask), opts )
  if err != nil {
    log.Error( "open boldtdb " + cfg.Database + " error: " + err.Error() )
    return nil, err
  }

  log.Info( "open boltdb " + cfg.Database )

  return &DB{ db: db }, nil
}


func (db *DB) Close() {
  db.db.Close()
}


func (db *DB) Get( bucket, key []byte ) []byte {

  var v []byte

  db.db.View( func (tx *bolt.Tx ) error {
     b := tx.Bucket(bucket)
     v = b.Get(key)
     return nil
  } )

  return v
}


func (db *DB) Set( bucket, key, value []byte ) {
  db.db.Update( func (tx *bolt.Tx) error {

    b, err := tx.CreateBucketIfNotExists(bucket)
    if err != nil {
        return err
    }

    return b.Put( key, value )
  } )
}


func (db *DB) Delete( bucket, key []byte ) {
    db.db.Update( func (tx *bolt.Tx) error {

    b, err := tx.CreateBucketIfNotExists(bucket)
    if err != nil {
        return err
    }

    return b.Delete( key )
  } )
}

