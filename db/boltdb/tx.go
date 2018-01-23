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

