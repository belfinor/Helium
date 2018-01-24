package proxy


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-01-24


import (
  "github.com/boltdb/bolt"
  "sync"
)


type DB struct {
  sync.Mutex
  db  *bolt.DB
  proxy *Proxy
  name  string
}


func (db *DB) Close() {
  db.Unlock()
}


func (db *DB ) View( f func (tx *bolt.Tx) error ) {
  db.db.View( f )
}


func (db *DB ) Update( f func (tx *bolt.Tx) error ) {
  db.db.Update( f )
}

