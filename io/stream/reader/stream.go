package stream


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-24


import (
  "bytes"
  "github.com/belfinor/Helium/pack"
  "os"
)


type Stream struct {
    data []byte
    file *os.File
}



type Reader interface {
  OnData( []byte )
}


func Read( filename string, r Reader ) {
  f, err := os.Open( filename )
  if err != nil {
    return
  }

  st := &Stream {
    file: f,
    data: make( []byte, 0, 4098 ),
  }

  st.read( r )
}


func (s *Stream) read(r Reader) {
 
  buf := make( []byte, 4098 )
 
  for {
    n, e := s.file.Read(buf)

    if e != nil {
        break
    }

    if n < 4098 {
      s.onRead( buf[0:n], r )
      break;
    } else {
      s.onRead( buf, r )
    }
  }
}


func (s *Stream) onRead( data []byte, r Reader ) {
  s.data = bytes.Join( [][]byte{  s.data, data }, []byte{} )
  size := int16(0)

  list := s.data

  for len(list) > 2 {
    if pack.Decode( list, &size ) != nil {
      break
    }
    size = size + 2
    if len(list) > int(size) {
      r.OnData( list[2:size] )
      list = list[size:]
    } else if len(list) == int(size) {
      r.OnData( list[2:] )
      list = []byte{}
    } else {
      break;
    }
  }

  if len(list) > 0 {
    s.data = list
  } else {
    s.data = []byte{}
  }
}

