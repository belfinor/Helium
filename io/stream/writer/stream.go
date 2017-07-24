package stream


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-24


import (
  "bytes"
  "github.com/belfinor/Helium/pack"
)


type Stream struct {
    data []byte
    writer *Writer
}


func NewWriter( cfg *WriterConfig ) *Stream {
  st := &Stream {
    data: make( []byte, 0, 4098 ),
    writer: InitWriter( cfg ),
  }

  return st
}


func (s *Stream) Add( src []byte ) {
  size := int16(len(src))
  s.Write( bytes.Join( [][]byte{ pack.Encode( size ), src }, nil ) )
}


func (s *Stream) Write( src []byte ) {
  s.data = bytes.Join( [][]byte{  s.data, src }, []byte{} )
  size := int16(0)
  
  list := s.data

  for len(list) > 2 {
    if pack.Decode( list, &size ) != nil {
      break
    }
    size = size + 2
    if len(list) > int(size) {
      s.writer.Push( list[:size] )
      list = list[size:]
    } else if len(list) == int(size) {
      s.writer.Push( list )
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


