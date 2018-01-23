package boltdb


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-01-23


import (
  "bytes"
  "github.com/boltdb/bolt"
)


type Bucket struct {
  b *bolt.Bucket
}


func (b *Bucket) Set( key, value []byte ) error {
  return b.b.Put( key, value )
}


func (b *Bucket) Get( key []byte ) []byte {
  return b.b.Get( key )
}


func (b *Bucket) Delete( key []byte ) error {
  return b.b.Delete( key )
}


func (b *Bucket) Range( from, to []byte, f FETCH_CALLBACK ) {

  c := b.b.Cursor()

  for k, v := c.Seek(from); k != nil && bytes.Compare(k, to) <= 0; k, v = c.Next() {

    if !f( k, v ) {
      break
    } 
  }
}


func (b *Bucket) Prefix( prefix []byte, f FETCH_CALLBACK ) {

  c := b.b.Cursor()

  for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {

    if !f( k, v ) {
      break
    } 
  }
}

