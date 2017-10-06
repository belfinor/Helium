package client


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-10-06


import (
  "testing"
)


func TestClient( t *testing.T) {

  www := New()

  res, err := www.Request( "GET", "https://morphs.ru", nil, nil )

  if err != nil {
    t.Fatal( "send request error" )
  }

  if res.Content == nil || len(res.Content) == 0 {
    t.Fatal( "content empty" )
  }
}

