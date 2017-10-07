package sociation


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-10-07


import (
  "testing"
)


func TestGet( t *testing.T ) {
  res := Get( "игра" )
  if res == nil || len(res) == 0 {
    t.Fatal( "not work" )
  }
}

