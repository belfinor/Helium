package slice


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-02-16


import (
  "math/rand"
  "reflect"
  "time"
)


func Shuffle( slice interface{} ) {

  rnd := rand.New( rand.NewSource( time.Now().UnixNano() ) )

  rv := reflect.ValueOf(slice)
  swap := reflect.Swapper(slice)
  length := rv.Len()

  for i := 0 ; i < length ; i++ {
    j := rnd.Intn(length)
    if i != j {
      swap(i, j)
    }
  }
}

