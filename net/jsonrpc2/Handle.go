package jsonrpc2


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2017-07-06


import (
    "encoding/json"
    "strings"
)


func Handle( in []byte ) []byte {
  
    text := strings.TrimSpace( string(in) )

    if text[0] == '{' {

        var rec Request
        var resp *Response

        if err := json.Unmarshal( []byte(text), &rec ) ; err != nil {
            return []byte("invalid request")
        } 
            
        resp = Exec( &rec )

        if resp == nil {
            return nil 
        }

        out, _ := json.Marshal( resp )
        return out

  } else if text[0] == '[' {

      recs := make( []Request, 0 )
      resp := make( []*Response, 0, 10 )

      if err := json.Unmarshal( []byte(text), &recs ) ; err != nil {
          return []byte("invalid request")
      }

      for _, rec := range recs {

          if r := Exec( &rec ) ; r != nil {
              resp = append( resp, r )
          }
      }

      out, _ := json.Marshal( resp )
      return out

  }

  return []byte("unsupported object type")
}

