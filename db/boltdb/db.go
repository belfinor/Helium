package boltdb


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2018-01-23


import (
  "bytes"
  "errors"
  "github.com/belfinor/Helium/log"
  "github.com/boltdb/bolt"
  "os"
  "time"
)


type DB struct {
  db *bolt.DB
}


type FETCH_CALLBACK func (key, value []byte) bool
type TRANSACTION_CALLBACK func (tx *Tx) error


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


func (db *DB) NextId( bucket, key, value []byte ) int64 {

  var id int64

  db.db.Update( func (tx *bolt.Tx) error {

    b, err := tx.CreateBucketIfNotExists(bucket)
    if err != nil {
        return err
    }

    i, err := b.NextSequence()
    id = int64(i)
    return err
  } )

  return id
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


func (db *DB) Prefix( bucket, prefix []byte, f FETCH_CALLBACK ) {

  db.db.View( func (tx *bolt.Tx) error {

    c := tx.Bucket(bucket).Cursor()

    for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {

      if !f( k, v ) {
        break
      } 

    }

    return nil

  } )

}


func (db *DB) Range( bucket, from, to []byte, f FETCH_CALLBACK ) {

  db.db.View( func (tx *bolt.Tx) error {

    c := tx.Bucket(bucket).Cursor()

    for k, v := c.Seek(from); k != nil && bytes.Compare(k, to) <= 0; k, v = c.Next() {

      if !f( k, v ) {
        break
      } 

    }

    return nil

  } )

}


func (db *DB) ForEach( bucket []byte, f FETCH_CALLBACK ) {

  db.db.View( func (tx *bolt.Tx) error {

    c := tx.Bucket(bucket)
    c.ForEach( func ( k, v []byte ) error {
      if !f(k,v ) {
        return errors.New( "stop iteration" )
      }
      return nil
    } )

    return nil

  } )
}


func (db *DB) Transaction( f TRANSACTION_CALLBACK ) {

  db.db.Update( func (tx *bolt.Tx) error {
    tr := &Tx{ tx: tx  }
    return f(tr)
  } )
}

