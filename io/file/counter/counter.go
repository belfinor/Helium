package counter


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-01-19


import (
  "github.com/belfinor/Helium/time/strftime"
  "io/ioutil"
  "strconv"
  "strings"
  "time"
)


type Counter struct {
  val  int64
  file string
}


func New( filename string, force bool ) ( *Counter, error ) {

  data, err := ioutil.ReadFile( filename )
  if err != nil {
    if force {
      return &Counter{ file: filename, val: 0 }, nil
    }
    return nil, err
  }

  c := &Counter{ file: filename, val: 0 }

  str := strings.TrimSpace( string(data) )
  
  c.val, err = strconv.ParseInt( str, 10, 64 )
  if err != nil {
    c.val = 0
  }

  return c, nil
  

}


func (c *Counter) Flush() {
  data := strconv.FormatInt( c.val, 10 )
  ioutil.WriteFile( c.file, []byte(data), 0644 )
}


func (c *Counter) Set( val int64 ) {
  c.val = val
  c.Flush()
}


func (c *Counter) Inc() {
  c.val++
  c.Flush()
}


func (c *Counter) SetDate( t time.Time ) {
  str := strftime.Format( "%Y%m%d", t )
  c.val, _ = strconv.ParseInt( str, 10, 64 )
  c.Flush()
}

