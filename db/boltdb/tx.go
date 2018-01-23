package boltdb


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-01-23


import (
  "github.com/boltdb/bolt"
)


type Tx struct {
  tx *bolt.Tx
}


func (tx *Tx) Bucket( name string ) ( *Bucket, error ) {
  b, err := tx.tx.CreateBucketIfNotExists( []byte(name) )
  if err != nil {
    return nil, err
  }

  bucket := &Bucket{ b: b }

  return bucket, nil
}


func (tx *Tx ) DeleteBucket( name string ) error {
  return tx.tx.DeleteBucket( []byte(name) ) 
}

