package filecache


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-13


import (
    "github.com/golang/groupcache"
    "io/ioutil"
)


var fCache *groupcache.Group


func Init( size int64 ) {
    if fCache == nil {
        fCache = groupcache.NewGroup( "filecache", size, groupcache.GetterFunc(
            func(ctx groupcache.Context, key string, dst groupcache.Sink) error {
                  data, err := ioutil.ReadFile( key )
                  if err != nil {
                      return err
                  }
                  dst.SetBytes(data)
                  return nil
            }),
        )
    }
}


func Get( filename string ) []byte {
    Init(10*1024*1024)
    var data []byte
    if fCache.Get( nil, filename , groupcache.AllocatingByteSliceSink( &data ) ) != nil {
      return nil
    }
    return data
}

