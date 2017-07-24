package stream


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-24


import (
  "fmt"
  "github.com/belfinor/Helium/log"
  "io/ioutil"
  "os"
  "time"
)


type Writer struct {
    Input chan []byte
    File  *os.File
    cfg   *WriterConfig
}


func InitWriter( cfg *WriterConfig ) *Writer {

  w := &Writer {
    Input: make( chan []byte, cfg.Buffer ),
    cfg: cfg,
  }

  w.openLog()

  log.Info( fmt.Sprintf( "storage current file=%s/%d", cfg.Path, cfg.LogId ) )
  log.Info( fmt.Sprintf( "storage buffer size=%d", cfg.Buffer ) )

  go w.Writer()

  return w
}


func (w *Writer) openLog() {

  file_name := fmt.Sprintf( "%s/%d", w.cfg.Path, w.cfg.LogId )
  var err error

  if w.File != nil {
    w.File.Close()
  }

  if w.File, err = os.OpenFile( file_name, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0664 ) ; err != nil {
    log.Error( err.Error() )
    panic(err)
  }

  ioutil.WriteFile( w.cfg.Index, []byte( fmt.Sprintf( "%d", w.cfg.LogId ) ), 664 )

  log.Info( "start log " + file_name )
}


func (w *Writer) rotate() {

  w.cfg.LogId++

  w.openLog()

  remove_name := fmt.Sprintf( "%s/%d", w.cfg.Path, w.cfg.LogId - w.cfg.Save )
  os.Remove( remove_name )
}


func (w *Writer) Push( data []byte ) {
  if data == nil ||  len(data) <= 2 {
    return  
  }

  block := make( []byte, len(data) )
  copy( block, data )

  w.Input <- block
}


func (w *Writer) Writer() {
  log.Info( "start storage writer" )

  start  := time.Now().Unix()
  last   := start

  period := w.cfg.Period

  for {
    select {
      case data := <- w.Input:
        if _, err := w.File.Write( data ) ; err != nil {
          panic(err)
        }
      case <- time.After( time.Second ):
    }

    last = time.Now().Unix()

    if last - start >= period {
      w.rotate()
      start = last
    }
  }
}

